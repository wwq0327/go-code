package main

import (
	"fmt"
)

func main() {
	var v = make(map[string]interface{}) // interface{} 让键值可以是任意类型的值
	v["name"] = "wyatt"
	v["age"] = 12

	for key, val := range v {
		fmt.Println(key, ":", val)
	}
}
