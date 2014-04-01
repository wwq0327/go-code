package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Printf("Slice: %s, Type: %T\n", s, s)
	fmt.Println(strings.Join(s, " "))

	a1 := map[int]string {1:"cccc"}
	a2 := []string{"1111", "2222"}

	fmt.Printf("%T:%v \n%T:%v\n",a1, a1, a2, a2) 
}
