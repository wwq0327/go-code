package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	URL = "http://www.xiaohuangji.com/ajax.php"
)

func main() {
	var input string
	for {
		fmt.Print("我说： ")
		fmt.Scanf("%s", &input)
		if input == "q" {
			os.Exit(1)
		} else {
			fmt.Printf("小黄说： %s\n", XiaoHuangJi(input))
			fmt.Println(strings.Repeat("-", 40))
		}
	}
}

func XiaoHuangJi(input string) string {
	data := url.Values{}
	data.Set("para", input)
	resp, err := http.PostForm(URL, data)
	if err != nil {
		log.Fatal("读取不出来，出问题啦！")
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		output, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("数据读不出来呀！")
		}

		return string(output)
	}
	return ""
}
