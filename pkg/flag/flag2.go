package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	boolPtr := flag.Bool("fork", false, "a bool")
	numbPtr := flag.Int("numb", 42, "an int")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("numb:", *numbPtr)

	fmt.Println("tail: ", flag.Args())

	if *wordPtr == "hello" {
		hello()
	}
}

func hello() {
	fmt.Println("hello")
}
