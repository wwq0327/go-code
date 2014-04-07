package main

import (
	"fmt"
	"net/url"
)

func main() {
	myurl := "http://www.google.com/index.html"
	u, _ := url.Parse(myurl)
	fmt.Println(u.Path)
}
