// 返回页面

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

	t,_ := template.ParseFiles("test.html")
	// 下方这种方式也行
	//	tneo := template.New("test.html")
	//	t.ParseFiles("test.html")


	t.Execute(w,"hello,everyone")
	// html嵌入信息
	// 无信息就nil

}

func handle2(w http.ResponseWriter, r *http.Request)  {

	// 用must处理err
	t := template.Must(template.ParseFiles("test.html"))
	t.Execute(w,"hello")

}
