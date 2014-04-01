package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Day())
	fmt.Println(time.Now().Day())
}
