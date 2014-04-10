package main

import (
	"fmt"
	"html"
	"strings"
)

func main() {
	s := html.EscapeString("<h2>hello world</h2>")
	fmt.Println(s)

	escapChars := `<>`

	i := strings.IndexAny("<b>hello world</b>", escapChars)

	fmt.Println(i, string(escapChars[i]))
}
