package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main()  {


	mux := httprouter.New()

	mux.GET("/helloworld/:id/:pwd",do)


	server := http.Server{
		Addr: ":8888",
		Handler: mux,
	}


	server.ListenAndServeTLS("cert.pem","key.pem")
	// 用crypto生成证书
	// https


}


func do(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	fmt.Fprintln(w,"this is do1111")
	fmt.Fprintf(w,"hello, %s, welcome\n",p.ByName("id"))
	fmt.Fprintf(w,"your pwd is %s\n",p.ByName("pwd"))

}
