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



	query := r.URL.Query()
	// 获取query

	fmt.Println(query.Get("hello"))
	// 如果key在query有多个，第一个
	fmt.Println(query["hello"])
	// 以slice形式返回所有值


	// query
	// query.Add(key,value)
	// query.Del(key)
	// query.Set(key,value)
	// query.Has(key)

}
