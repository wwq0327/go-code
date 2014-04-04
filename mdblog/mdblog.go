package main

import (
	"fmt"
	"os"
	"log"
	"time"
	"strings"
	"path/filepath"
	//"html/template"
	"io/ioutil"
	"github.com/russross/blackfriday"
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
	
	p := getPosts()
	fmt.Println(p)
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

// 获取所有.md文件信息
func getPosts() []MDBlog {
	a := []MDBlog{}
	files, _ := filepath.Glob("./contents/*")
	for _, f := range files {
		file := strings.Replace(f, "contents/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread,), "\n")
		title := string(lines[0])
		datetime := string(lines[1])
		content := strings.Join(lines[3:len(lines)], "\n")
		content = string(blackfriday.MarkdownCommon([]byte(content)))
		a = append(a, MDBlog{file, title, datetime, content})
	}
	return a
}


// 将所有的.md文件生成html文件
/*
func genHtml() {
	posts := getPosts()
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
}
*/	
