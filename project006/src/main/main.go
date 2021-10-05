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

	li := r.Header
	// header是map的子类

	fmt.Println(li)

	fmt.Println(li["Accept-Language"])
	// 返回数组
	fmt.Println(li.Get("Accept-Language"))
	// 返回一个 上述数组元素组成、由","连接

	// header方法
	// li.Add(key,value)
	// li.Del(key)
	// li.Set(key,value)
	// li.Clone()


}
