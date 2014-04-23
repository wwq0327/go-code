package main

import (
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	URL = "http://www.xiaohuangji.com/ajax.php"
)

var UserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1541.0 Safari/537.36"

func main() {
	var input string
	for {
		fmt.Print("我说： ")
		fmt.Scanf("%s", &input)
		if input == "q" {
			os.Exit(1)
		} else {
			ht, _ := HttpGet(&http.Client{}, URL, input, UserAgent)
			p, _ := ioutil.ReadAll(ht)
			fmt.Printf("小黄说： %s\n", string(p))
			fmt.Println(strings.Repeat("-", 40))
		}
	}
}

func HttpGet(client *http.Client, siteurl string, input string, header string) (io.ReadCloser, error) {
	v := url.Values{}
	v.Set("para", input)
	data := v.Encode()
	req, err := http.NewRequest("GET", siteurl, strings.NewReader(data))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	req.Header.Set("User-Agent", header)
	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	if response.StatusCode == 200 {
		return response.Body, nil
	}

	return nil, err
}
