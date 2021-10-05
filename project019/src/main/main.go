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
	li := make([]interface{},3)
	li[0],li[1],li[2] = 13,true,"glad to"
	t.Execute(w,li)

}
