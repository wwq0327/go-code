package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("w.txt")
	checkError(err)

	defer f.Close()

	f.WriteString("hello, world\n")

}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}
