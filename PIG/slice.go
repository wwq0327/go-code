package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Day())
	fmt.Println(time.Now().Day())

	re_str := fmt.Sprintf("<td class=\"(cur|ihbg)\">.+?<dt>%s</dt>(.+?)</dl>", "1号")
	fmt.Println(re_str)
}
