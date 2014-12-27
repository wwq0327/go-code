package main

import (
	"fmt"
)

type Humaner interface {
	Say()
}

type Human struct {
	Name string
}

func (h Human) Say() {
	fmt.Println(h.Name)
}

func main() {
	m := Human{"wyatt"}
	var h Humaner
	h = m
	h.Say()
}
