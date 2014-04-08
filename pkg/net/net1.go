package main

import (
	"fmt"
	"net"
)

func main() {
	a, _ := net.LookupHost("www.163.com")
	fmt.Printf("%v\n", a)
}
