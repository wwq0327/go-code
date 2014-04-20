// 自动创建一个可以生成jekyll源文件的代码

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	OUTPUT  = "./_posts"
	CONTENT = `---
layout: post
title:
description:
keywords:
---
`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <post title>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	title := os.Args[1]
	pn := PostName(title)
	WriteToFile(pn)
}

// 输入标题之后，接到当时日期后
func PostName(title string) string {
	t := time.Now().Format("2006-01-02")
	fileName := t + "-" + title + ".markdown"
	return fileName
}

// 将文件头内容写入到文件中
func WriteToFile(fileName string) {
	fp := filepath.Join(OUTPUT, fileName)
	f, err := os.Create(fp)
	defer f.Close()
	if err != nil {
		log.Fatal("文件打开错误！")
	}

	f.WriteString(CONTENT)

	fmt.Printf("文件位置：%s\n", fp)
}
