package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage: tuple <package-name> <type-name> <type>...")
		return
	}

	packageName := os.Args[1]
	typeName := os.Args[2]
	types := os.Args[3:]

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package " + packageName)
	buf.WriteString("\n\n")
	buf.WriteString("type " + typeName + " struct {\n")
	for i, t := range types {
		buf.WriteString("\t" + "V" + strconv.FormatInt(int64(i+1), 10) + " " + t + "\n")
	}
	buf.WriteString("}\n")

	buf.WriteString("\n")

	buf.WriteString("func New" + typeName + "(")
	for i, t := range types {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("v" + strconv.FormatInt(int64(i+1), 10) + " " + t)
	}
	buf.WriteString(") " + typeName + " {\n")
	buf.WriteString("\treturn " + typeName + "{\n")
	for i := range types {
		buf.WriteString("\t\t" + "V" + strconv.FormatInt(int64(i+1), 10) + ": v" + strconv.FormatInt(int64(i+1), 10) + ",\n")
	}
	buf.WriteString("\t}\n")
	buf.WriteString("}\n")

	rs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./" + strings.ToLower(typeName) + ".go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(rs)
	if err != nil {
		log.Fatal(err)
	}
}
