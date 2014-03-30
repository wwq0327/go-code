package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Hostname() (name string, err error)
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 
	}
	fmt.Printf("The hostname is: %s\n", name)

	// TempDir() string
	fmt.Printf("The temp dir is: %s\n", os.TempDir())
}
