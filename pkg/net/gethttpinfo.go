// get http server infomation by golang

package main

import (
	"fmt"
	"net/http"
)

const (
	URL = "http://golang.org"
)

func main() {
	httpInfo(URL)
}

// get http server info
func httpInfo(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer res.Body.Close()

	header := res.Header

	for k, v := range header {
		fmt.Printf("%s: %s\n", k, v[0])
	}

}
