package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2} // 未初始化的元素值为0
	b := [...]int{1, 2, 3, 4} // 通过初始化值来确定数组的长度
	c := [...]int{2: 100, 88:200} //使用索引号来初始化

	fmt.Println(a, b)
	fmt.Println(c)
}
