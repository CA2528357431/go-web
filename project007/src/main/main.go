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

	form := r.Form
	// 获取form+query
	// form中数据只支持x-www-form-urlencoded格式

	fmt.Println(form.Get("hello"))
	// 如果key在form、query都有，则优先返回form中的第一个

	fmt.Println(form["hello"])
	// 以数组返回form、query中的全部值

	// form
	// form.Add(key,value)
	// form.Del(key)
	// form.Set(key,value)
	// form.Has(key)

}
