package main

import (
	"fmt"
	"io/ioutil"
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


	files := multipartform.File
	// 文件map

	// 访问某key下的文件
	for _,file := range files["data"]{
		op,_ := file.Open()
		// 打开文件

		li,_ := ioutil.ReadAll(op)
		// 读取数据

		str := string(li)
	
		fmt.Println(str)
	}
}
