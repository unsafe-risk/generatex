package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: helloworld <package-name>")
		return
	}
	packageName := os.Args[1]

	f, err := os.Create("./helloworld.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	buf.WriteString("package " + packageName)
	buf.WriteString("\n\n")
	buf.WriteString("func HelloWorld() string {\n")
	buf.WriteString("\treturn \"Hello World\"\n")
	buf.WriteString("}\n")

	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
