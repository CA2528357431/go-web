package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/helloworld",line2_3(line1_2(do1)))


	http.ListenAndServe(":8888",nil)


}

// 处理器串联

func do1(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"this is do1111")

}

func line1_2(h1 http.HandlerFunc)http.HandlerFunc{
	neo := func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintln(w,"this is do2222")
		// 新代码部分

		// 串联的代码复用
		h1(w,r)
	}
	return neo
}

func line2_3(h2 http.HandlerFunc)http.HandlerFunc{
	neo := func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintln(w,"this is do333")
		h2(w,r)
	}
	return neo
}