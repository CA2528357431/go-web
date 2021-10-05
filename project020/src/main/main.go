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
	t.Execute(w,"false and true")

}

// with end之间的 . 改变为with后的值
// else只有在with为空或为0时执行
