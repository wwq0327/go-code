package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"html/template"
)

const (
	UPLOAD_DIR = "./uploads"
)

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {		
		if err := renderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}


// view the image
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id") // 获取图片id
	imagePath := UPLOAD_DIR + "/" + imageId // 重组图片访问地址

	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath) // 从路径读到文件，并将信息输入到客户端
}

//判断文件是否真的存在
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// 表出所有可以访问的图片文件名
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	locals := make(map[string]interface{})
	images := []string{}

	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	
	locals["images"] = images
	
	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// 模版加载及渲染函数
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	t, err := template.ParseFiles(tmpl+".html")
	if err != nil {
		return
	}
	err = t.Execute(w, locals)
	return
}
