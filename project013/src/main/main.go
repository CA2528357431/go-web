package main

import (
	"net/http"
)

func main()  {

	http.HandleFunc("/helloworld1",handle1)

	http.HandleFunc("/helloworld2",handle2)

	http.HandleFunc("/helloworld3",handle3)
	// 通过此来处理对未存在功能的访问

	http.HandleFunc("/helloworld4",handle4)
	// 重定向

	http.ListenAndServe(":8888",nil)

}

func handle1(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello, this is 11111"))
}
func handle2(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello, this is 22222"))
}
func handle3(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(501)
	// 写入状态码
	// 通过修改状态吗可以改变传递的信息
	// 具体参见状态码501
}
func handle4(w http.ResponseWriter, r *http.Request)  {

	header := w.Header()
	// 获取w响应头，区分于上文的r请求头

	header.Set("Location","https://cn.bing.com")

	w.WriteHeader(302)
}