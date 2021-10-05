package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main()  {


	// httprouter
	mux := httprouter.New()

	mux.GET("/helloworld/:id/:pwd",do)
	// 用 :xx 传参

	// 只支持精准匹配
	// /helloworld/123就无法匹配到/helloworld/:id/:pwd


	http.ListenAndServe(":8888",mux)


}

// 处理器串联

func do(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	fmt.Fprintln(w,"this is do1111")
	fmt.Fprintf(w,"hello, %s, welcome\n",p.ByName("id"))
	fmt.Fprintf(w,"your pwd is %s\n",p.ByName("pwd"))
	// ByName参数就是上文的xx

}
