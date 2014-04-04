package main

import "fmt"
import "strconv"
import "net/url"
import "log"
import "time"


func main() {
	var a int = 65
	b := strconv.Itoa(a)
	fmt.Printf("%s\n", b)

	ss :="hello, world"
	fmt.Println([]byte(ss))

	u, err := url.Parse("https://google.com/search?q=golang")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(u.Scheme, u.Host, u.Query())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
