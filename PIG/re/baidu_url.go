package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	//"strings"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	fmt.Printf("%T\n", body)
	//src := string(body)
	//fmt.Printf("%T\n", src)

	re := regexp.MustCompile(`<a href="(?P<url>\S+)">.+</a>`)
	all := re.FindAll(body, -1)
	
	for _, v := range all {
		fmt.Println(string(v))
	}

}
