package http

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var examplePrefix = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>
`

func TestHttpGet(t *testing.T) {
	// 200
	rc, err := HttpGet(&http.Client{}, "http://example.com", nil)
	if err != nil {
		t.Fatalf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}

	p, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}

	s := string(p)
	if !strings.HasPrefix(s, examplePrefix) {
		t.Errorf("HttpGet:\n Expect => %v\n Got => %s\n", nil, err)
	}
}

func BenchmarkHttpGet(b *testing.B) {
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		HttpGet(client, "http://example.com", nil)
	}
}
