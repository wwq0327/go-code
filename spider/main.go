package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
/* 简易爬虫，读取一个页面 */

// 指定读取的http地址
const URL = "http://golang.org/doc/tos.html"

// 读取内容，返回内容和状态码
func Content(url string) (content string, status int) {
	res, err := http.Get(url) 
	if err != nil {
		fmt.Printf("%s\n", err)
		status = -100
		return 
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		status = -200
		return 
	}
	
	status = res.StatusCode
	content = string(data)
	return 
}

func main() {
	html, status := Content(URL)
	if status != http.StatusOK {
		fmt.Printf("Error code: %d\n", status)
		return 
	}
	fmt.Printf("%s\n", html)
}

