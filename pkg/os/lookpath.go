// 寻找文件的path
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
		return
	}

	lookpath(os.Args[1])
}

func lookpath(arg string) {
	path, err := exec.LookPath(arg)
	if err != nil {
		fmt.Printf("Error: %s\n", path)
		return
	}

	fmt.Printf("%s is available at %s\n", arg, path)
}
