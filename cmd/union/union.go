package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		log.Println("Usage: monad <package-name> <type-name> <types>...")
		return
	}

	packageName := os.Args[1]
	typeName := os.Args[2]
	types := os.Args[3:]

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package " + packageName)
	buf.WriteString("\n\n")
	buf.WriteString("type " + typeName + " struct {\n")
	buf.WriteString("\tvalue any\n")
	buf.WriteString("}\n\n")

	for _, t := range types {
		n := strings.ToUpper(t[:1]) + t[1:]
		buf.WriteString("func (m " + typeName + ") Is" + n + "() bool {\n")
		buf.WriteString("\t_, ok := m.value.(" + t + ")\n")
		buf.WriteString("\treturn ok\n")
		buf.WriteString("}\n\n")

		buf.WriteString("func (m " + typeName + ") As" + n + "() (" + t + ", bool) {\n")
		buf.WriteString("\tv, ok := m.value.(" + t + ")\n")
		buf.WriteString("\treturn v, ok\n")
		buf.WriteString("}\n\n")

		buf.WriteString("func (m " + typeName + ") Must" + n + "() " + t + " {\n")
		buf.WriteString("\tv, ok := m.value.(" + t + ")\n")
		buf.WriteString("\tif !ok {\n")
		buf.WriteString("\t\tpanic(\"" + t + " is not the type of the value\")\n")
		buf.WriteString("\t}\n")
		buf.WriteString("\treturn v\n")
		buf.WriteString("}\n\n")

		buf.WriteString("func " + typeName + "Of" + n + "(v " + t + ") " + typeName + " {\n")
		buf.WriteString("\treturn " + typeName + "{\n")
		buf.WriteString("\t\tvalue: v,\n")
		buf.WriteString("\t}\n")
		buf.WriteString("}\n\n")
	}

	rs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(strings.ToLower(typeName) + ".go")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(rs)
	if err != nil {
		log.Fatal(err)
	}
}
