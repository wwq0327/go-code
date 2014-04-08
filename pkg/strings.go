package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Contains(s, substr string) bool
	// s 需要判断的主串
	// 包含的子串

	fmt.Println(strings.Contains("wyattwang", "tt")) // true
	fmt.Println(strings.Contains("wyattwang", "yy")) // false

	// func ContainsAny(s, chars string) bool
	// 判断s中是否包含chars中的任意字符

	fmt.Println(strings.ContainsAny("wyattwang", "wiiic")) // true

	// 字符串s和t比较，它们在全部小写的情况下，采用UTF8编码的底层的unicode是否一致
	// 这个还较有用的
	fmt.Println(strings.EqualFold("Go", "go")) //true
}
