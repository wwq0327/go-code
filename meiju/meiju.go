package main

import (
	"fmt"
	//"os"
	"net/http"
	"io/ioutil"
	"time"
	"strconv"
	"strings"
	"regexp"
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

// 打印当天所在天的字符
func get_day_str() string {
	day := time.Now().Day()
	day_str := strconv.Itoa(day) + "号" // 取当天号数
	return day_str
}

// 获取当天的美剧时间表HTML
func get_day_html(html string, day string) []string {
	re_str := fmt.Sprintf("<td class=\"(cur|ihbg)\">.+?<dt>%s</dt>(.+?)</dl>", day)
	//fmt.Println(re_str)
	re, _ := regexp.Compile(re_str)
	src := re.FindAllString(html, -1)
	return src
}


// 打印表头
func print_str() {
	fmt.Printf("%s\n", strings.Repeat("=", 72))
}

func main() {
	day := get_day_str()
	fmt.Printf("%s 的美剧更新情况为:\n", day)
	print_str()
	html, status := get_content(url)
	if status != http.StatusOK {
		fmt.Printf("Error code: %d\n", status)
		return 
	}
	fmt.Printf("%T\n",html)
	day_html := get_day_html(html, day)
	fmt.Println(day_html)

}