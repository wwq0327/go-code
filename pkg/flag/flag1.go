package main

import (
	"flag"
)

var flagName string

func init() {
	flag.StringVar(&flagName, "name", "default name", "help name message.")
	flag.Parse()
}

func main() {
	flag.PrintDefaults()
}
