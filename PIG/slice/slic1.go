package main

import (
	"fmt"
)

func main() {
	s := []string{"A", "B", "C"}
	t := []string{"D", "E", "F"}
	fmt.Printf("%v -- %v\n", s, t)
	
	// 添加一个切片时，需要在切片名后有...
	s = append(s, t...)
	fmt.Printf("s:%v-- t:%v\n", s, t)
	t = append(t, "h", "k")
	fmt.Printf("s:%v-- t:%v\n", s, t)
}
