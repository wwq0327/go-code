package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "www.baidu.com")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}
