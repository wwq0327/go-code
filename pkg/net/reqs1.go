package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	s := reqEncode()
	fmt.Println(s)
}

func reqEncode() string {
	client := &http.Client{}
	v := url.Values{}
	v.Set("wd", "golang")
	data := v.Encode()
	fmt.Println(data)
	reqs, err := http.NewRequest("POST", "http://baidu.com", strings.NewReader(data))
	// fmt.Println(reqs)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	// reqs.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	fmt.Println(reqs.Header)
	response, err := client.Do(reqs)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	if response.StatusCode == 200 {
		var body []byte
		body, _ = ioutil.ReadAll(response.Body)
		return string(body)
	}

	return ""

}
