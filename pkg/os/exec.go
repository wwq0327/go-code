package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arg := []string{"Hello", "world"}
	cmd := exec.Command("echo", arg...)
	output, _ := cmd.Output()

	// checkError(err)

	fmt.Printf("The output is: %s\n%+v\n", output, cmd)

}
