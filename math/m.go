package main

import (
	"fmt"
)

func main() {
	num := []int{1, -2, 0}
	for _, x := range num {
		fmt.Println(Abs(float32(x)))
	}
}

// 求绝对值
func Abs(d float32) float32 {
	if d >= 0 {
		return d
	} else {
		return -d 
	}
}
