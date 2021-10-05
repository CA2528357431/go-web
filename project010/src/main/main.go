package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/helloworld",do)

	http.ListenAndServe(":8888",nil)

}

func do(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello\nworld\n")

	r.ParseMultipartForm(1024)
	// 必须解析multipart form
	// 参数表示解析出的数据量

	multipartform := r.MultipartForm
	// 获取postform
	// 只支持form-data格式

	value := multipartform.Value
	// map储存键值对

	fmt.Println(value["hello"])
	// 以数组返回form中的全部值

}
