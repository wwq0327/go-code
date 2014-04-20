package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/raw", rawHandler)
	http.HandleFunc("/html", htmlHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var listMd []string
	for _, fileInfo := range fileInfoArr {
		filename := fileInfo.Name()
		fileExt := filepath.Ext(filename)
		if fileExt == ".md" || fileExt == ".markdown" {
			listMd = append(listMd, filename)
		}
	}

	t := template.New("index template")
	t, _ = template.ParseFiles("index.html")
	t.Execute(w, listMd)
}

// 直接显示出源文件内容
func rawHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("name")
	mdpath := "./" + filename

	w.Header().Set("Content-Type", "text")
	http.ServeFile(w, r, mdpath)
}

// 显示通过Markdown格式化后的内容
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("name")
	mdpath := "./" + filename
	md, err := ioutil.ReadFile(mdpath)
	if err != nil {
		log.Fatal("读取错误")
	}
	content := template.HTML(string(blackfriday.MarkdownCommon(md)))

	t := template.New("html template")
	t, _ = template.ParseFiles("md.html")

	t.Execute(w, content)
}
