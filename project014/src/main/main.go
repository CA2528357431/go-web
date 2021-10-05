package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {

	http.HandleFunc("/helloworld1", send)
	http.HandleFunc("/helloworld2", get)

	http.ListenAndServe(":8888",nil)

}

func send(w http.ResponseWriter, r *http.Request)  {
	c := http.Cookie{
		Name:       "Mycookie",
		// 只能包含字母或者数字
		// 不能重复
		Value:      "hello world",
		Expires:    time.Now().Add(time.Second*2000),
		// 如果为空，则 Cookie 就是 Session Cookie，这个 Cookie 会随着浏览器的关闭而销毁
		// 否则这个 Cookie 就是持久 Cookie，直到过期时间后才会销毁
		// MaxAge: xx
		// 创建xx后过期,推荐还是用expire
		HttpOnly:   true,
	}
	w.Header().Set("Set-Cookie",c.String())
	// 添加cookie

	cc := http.Cookie{
		Name:       "Fastcookie",
		Value:      "see you world",
		Expires:    time.Now().Add(time.Second*2000),
		HttpOnly:   true,
	}

	http.SetCookie(w,&cc)
	// 快捷添加cookie

	fmt.Println("send")
}

func get(w http.ResponseWriter, r *http.Request)  {
	c,_ := r.Cookie("Joker")
	fmt.Println(c)
	// 获取指定cookie

	fmt.Println(r.Cookies())
	// 获取全部cookie
}

