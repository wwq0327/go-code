package main

import (
	"fmt"
	"os/exec"
)

// found file path, else print error message
func main() {
	f, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println("not install ls")
	}
	fmt.Println(f)

	argv := []string{"-lh"} // 参数，一个字符型slice
	c := exec.Command("ls", argv...)
	d, _ := c.Output()
	fmt.Println(string(d))
}
