package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"path/filepath"
)

func main() {
	s := []byte("##hello")
	output := blackfriday.MarkdownBasic(s)
	fmt.Printf("%s\n", output)

	files, _ := filepath.Glob("./contents/*")
	fmt.Println(files)
}
