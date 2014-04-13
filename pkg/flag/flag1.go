package main

import (
	"flag"
	"fmt"
)

var flagName string

func init() {
	flag.StringVar(&flagName, "name", "wyatt", "help name message.")
	flag.Parse()
}

func main() {

	fmt.Println(flag.Args())
}

func Hello() {
	fmt.Println("hello, world")
}
