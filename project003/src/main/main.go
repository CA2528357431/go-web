package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/do1",do1)

	http.HandleFunc("/do2/",do2)
	// 最小惊讶原则
	// URL最前一层搜索失败则转向上一级寻找是否有匹配的handler
	// 在url后加一个"/"


	http.ListenAndServe(":8888",nil)


}

// 处理器串联

func do1(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"this is do1111")

}
func do2(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"this is do2222")

}
