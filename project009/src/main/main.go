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

	r.ParseForm()
	// 必须解析form

	postform := r.PostForm
	// 获取postform
	// postform中数据只支持x-www-form-urlencoded格式

	fmt.Println(postform.Get("hello"))
	// 返回form中的第一个

	fmt.Println(postform["hello"])
	// 以数组返回form中的全部值

	// postform
	// postform.Add(key,value)
	// postform.Del(key)
	// postform.Set(key,value)
	// postform.Has(key)

}
