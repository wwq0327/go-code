package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	html = `<html>
	<head>
	<title>hello</title>
	</head>
	<body>
	<h1>hello, world!</h1>
	</body></html>`
)

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, html)
}
