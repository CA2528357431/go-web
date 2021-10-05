package main

import (
	"encoding/base64"
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
		Name:       "Flash",
		Value:      base64.URLEncoding.EncodeToString([]byte("instead of a flash")),
		// 真正的value base64加密
	}
	// 签发cookie签证

	http.SetCookie(w,&c)

	fmt.Fprintln(w,"set the passport")

	fmt.Println("set")
}

func get(w http.ResponseWriter, r *http.Request)  {
	c,err := r.Cookie("Flash")
	if err!=nil{
		if err==http.ErrNoCookie{
			fmt.Fprintln(w,"no such right")
			// 无cookie签证

			fmt.Println("fail")
		}
	}else {
		cc := http.Cookie{
			Name:       "Flash",
			Expires: time.Now().Add(time.Second*(-1)),
			// 真正的value base64加密
		}
		http.SetCookie(w,&cc)
		// 更新cookie签证，使就cookie失效

		value,_ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w,string(value))



		fmt.Println("succeed")

	}
}

