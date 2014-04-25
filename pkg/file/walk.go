package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/Users/wyatt/Downloads", walkFunc)
}

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}
