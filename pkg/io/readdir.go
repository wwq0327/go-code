package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
)

const (
	BASE = "./"
)

func main() {
	fp := filepath.Join(BASE, "test")
	fmt.Println(fp)
	fmt.Println(filepath.Abs(fp))

	fs, err := ioutil.ReadDir(fp)
	if err != nil {
		panic(err)
		return
	}

	for _, fileInfo := range fs {
		fmt.Println(fileInfo.Name()) // 文件名
		fmt.Println(filepath.Join(BASE, fileInfo.Name()))

		fmt.Println(path.Ext(fileInfo.Name())) // 读取扩展名
	}

}
