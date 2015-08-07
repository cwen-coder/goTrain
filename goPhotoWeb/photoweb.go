package main

import (
	"html/template"
	"io/ioutil"
	"os"
	//"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"runtime/debug"
	"strings"
)

const (
	UPLOAD_DIR   = "./uploads"
	HTML_START   = "<html>"
	HTML_END     = "</html>"
	TEMPLATE_DIR = "./views"
	ListDir      = 0x0001
)

var templates map[string]*template.Template = make(map[string]*template.Template)

//templates = make(map[string]*template.Template)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//t, err := template.ParseFiles("upload.html")
		err := renderHtml(w, "upload", nil)
		check(err)
		//t.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	//var listHtml string
	locals := make(map[string]interface{})
	images := []string{}

	for _, fileInfo := range fileInfoArr {
		//imgid := fileInfo.Name()
		//listHtml += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	//t, err := template.ParseFiles("list.html")
	//io.WriteString(w, HTML_START+"<ol>"+listHtml+"</ol>"+HTML_END)
	err = renderHtml(w, "list", locals)
	check(err)
	//t.Execute(w, locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
	/*t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		return err
	}
	err = t.Execute(w, locals)*/
	err := templates[tmpl].Execute(w, locals)
	return err
}

func init() {
	/*for _, tmpl := range []string{"upload", "list"} {
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}*/
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		arr := strings.Split(templateName, ".")
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[arr[0]] = t
	}
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/", safeHandler(listHandler))
	err := http.ListenAndServe(":5050", mux)
	if err != nil {
		log.Fatal("ListenAndServer:", err.Error())
	}
}
