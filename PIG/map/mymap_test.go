package mymap

import (
	// "fmt"
	"testing"
)

func TestMymap(t *testing.T) {
	slice := []int{9, 2, 9, 3, 8, 3}
	uniqueSlice := []int{9, 2, 3, 8}
	s := UniqueInts(slice)
	for k, v := range s {
		if v == uniqueSlice[k] {
			t.Log("pass")
		} else {
			t.Error("测试没通过")
		}
	}
}
