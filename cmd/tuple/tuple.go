package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: tuple <package-name> <type>...")
		return
	}

	packageName := os.Args[1]
	types := os.Args[2:]

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package " + packageName)
	buf.WriteString("\n\n")
	buf.WriteString("type Tuple struct {\n")
	for i, t := range types {
		buf.WriteString("\t" + "V" + strconv.FormatInt(int64(i+1), 10) + " " + t + "\n")
	}
	buf.WriteString("}\n")

	buf.WriteString("\n")

	buf.WriteString("func NewTuple(")
	for i, t := range types {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("v" + strconv.FormatInt(int64(i+1), 10) + " " + t)
	}
	buf.WriteString(") Tuple {\n")
	buf.WriteString("\treturn Tuple{\n")
	for i := range types {
		buf.WriteString("\t\t" + "V" + strconv.FormatInt(int64(i+1), 10) + ": v" + strconv.FormatInt(int64(i+1), 10) + ",\n")
	}
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")

	f, err := os.Create("./tuple.go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
