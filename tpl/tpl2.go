package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	t := template.New("html")
	t.Execute(os.Stdout, "<h2>Hello, world</h2>")
}
