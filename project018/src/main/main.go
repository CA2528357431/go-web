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

	x := 1
	t.Execute(w,x>5)

}

