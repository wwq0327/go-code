package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("Fieldname exemple")
	t, _ = t.Parse("hello, <h1>{{ .UserName}}!</h1>")
	p := Person{"Wyatt"}
	t.Execute(os.Stdout, p)

}
