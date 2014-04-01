package main

import (
	"fmt"
	//"os"
	"net/http"
	"io/ioutil"
)

const url = "http://www.yyets.com/tv/schedule"

// 读取指定页面的所有HTML内容
func get_content(url string)(content string, status int) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err)
		status = -100
		return
	}
	
	defer res.Body.Close() // 关闭读取
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		status = -200
		return 
	}

	status = res.StatusCode // 读取状态码
	content = string(data)  // 将[]byte转换成string

	return content, status
}

func main() {
	html, status := get_content(url)
	if status != http.StatusOK {
		fmt.Printf("Error code: %d\n", status)
		return 
	}
	fmt.Printf("%s\n", html)
}
