package main

import (
	"fmt"
	"os"
	"log"
	"time"
	//"strings"
	"path/filepath"
)

const (
	MD_DIR = "./contents"  // markdown 源文件目录
    OUTPUT = "./output"    // 输出文件目录
)

type MDBlog struct {
	filename string   // 文件名
	title string
	datetime string   //发布时间 
	content string    // 内容
}

func main() {
	//fmt.Println(os.Args)

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		fmt.Printf("Usage: %s -new <filename> : create new blog file \n\t -gen generate html\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}
 
	if os.Args[1] == "-new" {
		blog := MDBlog{filename:os.Args[2], title:os.Args[2]}
		blog.createNewPost()
	}

}

// 创建一个新的文件
func (blog MDBlog) createNewPost() error {
	md_file_path := MD_DIR + "/" + blog.filename + ".md"

	if exists := isExists(md_file_path); exists {
		log.Fatal("file is exists!")
		os.Exit(0)
	}

	fmt.Printf("new file: %s\n", md_file_path)
	file, err := os.Create(md_file_path)
	if err != nil {
		log.Fatal("Failed to create the output file", err)
		return err
	}

	defer file.Close()
	datetime := time.Now().Format("2006-01-02 15:04:05")
	file.WriteString("Title: "+ blog.title + "\nDate: "+datetime+"\n\n")
	
	return nil
}

//判断文件是否真的存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
