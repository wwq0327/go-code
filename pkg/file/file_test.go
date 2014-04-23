package file

import (
	"testing"
)

func TestBasename(t *testing.T) {
	if name := Basename("./foo/file.go"); name == "file.go" {
		t.Log("test pass")
	} else {
		t.Error("test not pass")
	}
}

func TestDirname(t *testing.T) {
	if name := Dirname("/a/b/c"); name == "/a/b" {
		t.Log("test pass")
	} else {
		t.Error("test not pass")
	}
}

func BenchmarkBasename(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basename("file.go")
	}
}

func BenchmarkDirname(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dirname("/a/b/c")
	}
}
