package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/helloworld1",handle1)
	http.HandleFunc("/helloworld2",handle2)
	http.ListenAndServe(":8888",nil)

}

func handle1(w http.ResponseWriter, r *http.Request)  {

	// 同时解析多个html
	t := template.Must(template.ParseFiles("test1.html","test2.html"))
	// 则t是几个html集合

	t.Execute(w,"hello,everyone")
	// 执行第一条

}

func handle2(w http.ResponseWriter, r *http.Request)  {

	t := template.Must(template.ParseFiles("test1.html","test2.html"))

	t.ExecuteTemplate(w,"test2.html","hello,everyone")
	// 指定执行
	// 若不存在则返回空值



}
