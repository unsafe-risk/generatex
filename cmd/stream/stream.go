package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: stream <package-name> <types>...")
		return
	}

	packageName := os.Args[1]
	types := os.Args[2:]

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package " + packageName)
	buf.WriteString("\n\n")
	buf.WriteString("type Stream struct {\n")
	for i, t := range types[:len(types)-1] {
		buf.WriteString(fmt.Sprintf("\tF%d func(%s) (%s, error)\n", i+1, t, types[i+1]))
	}
	buf.WriteString("}\n")

	buf.WriteString("\n")

	buf.WriteString("func NewStream(")
	for i, t := range types[:len(types)-1] {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("f%d func(%s) (%s, error)", i+1, t, types[i+1]))
	}
	buf.WriteString(") *Stream {\n")
	buf.WriteString("\treturn &Stream{\n")
	for i := range types[:len(types)-1] {
		buf.WriteString(fmt.Sprintf("\t\tF%d: f%d,\n", i+1, i+1))
	}
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")

	buf.WriteString("\n")

	buf.WriteString("func (s *Stream) Run(")
	buf.WriteString(fmt.Sprintf("init %s", types[0]))
	buf.WriteString(") (")
	buf.WriteString(fmt.Sprintf("result %s", types[len(types)-1]))
	buf.WriteString(", err error) {\n")
	buf.WriteString("\trs1 := init\n")
	for i := range types[:len(types)-1] {
		buf.WriteString(fmt.Sprintf("\trs%d, err := s.F%d(rs%d)\n", i+2, i+1, i+1))
		buf.WriteString("\tif err != nil {\n")
		buf.WriteString("\t\treturn result, err\n")
		buf.WriteString("\t}\n\n")
	}
	buf.WriteString(fmt.Sprintf("\treturn rs%d, nil\n", len(types)))
	buf.WriteString("}\n")

	rs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./stream.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(rs)
	if err != nil {
		log.Fatal(err)
	}
}
