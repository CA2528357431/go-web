package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {

	http.HandleFunc("/helloworld1", handle1)
	http.HandleFunc("/helloworld2", handle2)

	http.ListenAndServe(":8888",nil)

}

func readData()[]byte  {
	file := "./test.html"
	li,_ := ioutil.ReadFile(file)
	return li
}
func handle1(w http.ResponseWriter, r *http.Request)  {
	// html

	// w.Header().Set("Content-Type","text/html")
	// 尽管没有手动设置数据类型，但可以自动检测前512字节分析类型


	li := readData()
	w.Write(li)

	r.ParseForm()
	form := r.Form
	fmt.Printf("username: %s\npassword: %s\n",form.Get("UserName"),form.Get("Password"))
}


func handle2(w http.ResponseWriter, r *http.Request)  {
	// json

	// w.Header().Set("Content-Type","application/json")

	type per struct {
		Age  int `json:"age"`
		Name string `json:"name"`
	}
	p := per{
		Age:  18,
		Name: "caoan",
	}

	li,_ := json.Marshal(p)

	w.Write(li)
}