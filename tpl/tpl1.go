package main

import (
	"html/template"
	"os"
	"fmt"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("Fieldname exemple")
	t, _ = t.Parse("hello, {{ .UserName}}!")
	p := Person{"Wyatt"}
	t.Execute(os.Stdout, p)

}

