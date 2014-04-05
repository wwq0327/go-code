package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MD_DIR = "./contents" // markdown 源文件目录
	OUTPUT = "./output"   // 输出文件目录
)

type MDBlog struct {
	Filename string // 文件名
	Title    string
	Datetime string //发布时间
	Content  string // 内容
}

func main() {
	//fmt.Println(os.Args)

	if len(os.Args) == 1 || os.Args[1] == "-h" {
		fmt.Printf("Usage: %s -new <filename> : create new blog file \n\t -gen generate html\n", filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	if os.Args[1] == "-new" {
		blog := MDBlog{Filename: os.Args[2], Title: os.Args[2]}
		blog.createNewPost()
	}

	if os.Args[1] == "-gen" {
		genHtml()
	}

}

// 创建一个新的文件
func (blog MDBlog) createNewPost() error {
	md_file_path := MD_DIR + "/" + blog.Filename + ".md"

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
	file.WriteString("Title: " + blog.Title + "\nDate: " + datetime + "\n\n")

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
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		datetime := string(lines[1])
		content := strings.Join(lines[3:len(lines)], "\n")
		content = string(blackfriday.MarkdownCommon([]byte(content)))
		a = append(a, MDBlog{file, title, datetime, content})
	}
	return a
}

// 将所有的.md文件生成html文件
func genHtml() {
	file, err := os.Create(OUTPUT + "/" + "index.html")
	if err != nil {
		fmt.Print(err)
	}
	posts := getPosts()
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	fmt.Println(posts)
	t.Execute(file, posts)
}
