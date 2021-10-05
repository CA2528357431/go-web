package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/helloworld",handle)
	http.ListenAndServe(":8888",nil)

}

func handle(w http.ResponseWriter, r *http.Request)  {

	t := template.Must(template.ParseFiles("test.html"))
	t.ExecuteTemplate(w,"main","promised")

}

// 区分与template 不需要多文件
