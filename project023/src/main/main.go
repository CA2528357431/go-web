package main

import (
	"net/http"
	"text/template"
	"time"
)

func main() {

	http.HandleFunc("/helloworld",handle)
	http.ListenAndServe(":8888",nil)

}

func handle(w http.ResponseWriter, r *http.Request)  {

	funcmap := template.FuncMap{
		"time":format,
		"timeadd":formatadd,
	}

	t := template.New("test.html").Funcs(funcmap)

	t,_ = t.ParseFiles("test.html")

	t.Execute(w,time.Now())

}

func format (t time.Time)string{
	str := t.Format("2006-01-02 15:04:05")
	return str
}
func formatadd (t time.Time,xx int)string{
	str := t.Add(time.Hour*time.Duration(xx)).Format("2006-01-02 15:04:05")
	return str
}

// {a|b|c}
// 将a的值传参给b，b求出b传给c，求出c显示


// 根据 . 的位置可以自动转型   有效防止注入型攻击  P119

