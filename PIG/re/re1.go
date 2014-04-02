package main

import (
	"fmt"
	"regexp"
)

func myFindAllString(re_string string, text string) {
	reg := regexp.MustCompile(re_string)
	fmt.Printf("%q\n",  reg.FindAllString(text, -1))
}

func main() {
	hello := "Hello, world! 世界, 你好, Go"
	
	myFindAllString(`\w+`, hello)

	//查找连续小写字母
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找连续非小写字母
	reg = regexp.MustCompile(`[^a-z]+`)
	fmt.Println(reg.FindAllString(hello, -1))

	// 查找连续单词
	reg = regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找连续的非单词字母、非空白字符
	reg = regexp.MustCompile(`[^\w\s]+`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找连续汉字
	reg = regexp.MustCompile(`[\p{Han}]+`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找连续非汉字
	reg = regexp.MustCompile(`[\P{Han}]+`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找以H开头，以空格结尾的字符串
	reg = regexp.MustCompile(`^H.*\s`)
	fmt.Println(reg.FindAllString(hello, -1))

	//查找行首以H开头，以空白结尾的字符串（非贪婪模式）
	reg = regexp.MustCompile(`(?U)^H.*\s`)
	fmt.Println(reg.FindAllString(hello, -1))

	// 查找以hello开头，忽略大小写，以Go结尾的字符串
	myFindAllString(`(?i:^hello).*Go`, hello)
}
