package main

import (
	"fmt"
	"net/http"
)



func mynet1() {

	http.HandleFunc("/helloworld",do)

	// 此时为一个正常的默认mux
	http.ListenAndServe(":8888",nil)
}

func do(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world")
	fmt.Fprintln(w,r.RemoteAddr)
}

func mynet2() {

	m := mux{}

	// 此时将一个处理器当成mux
	http.ListenAndServe(":9999",&m)

	// 发现无论什么url都是一个处理器
}

type mux struct {

}

func (m mux)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world")
	fmt.Fprintln(w,r.RemoteAddr)
}

// mux也是处理器，但其的作用是将网站目录于处理器匹配
