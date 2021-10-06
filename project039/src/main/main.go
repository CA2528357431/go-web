// 数据序列化为json

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	a := struct {
		Name string
		Age int
		School []string
	}{
		Name:   "ca",
		Age:    18,
		School: []string{"hust","hit"},
	}
	// 记得把struct的参数变成public

	data, err := json.Marshal(a)
	if err!=nil{
		println(err)
	}
	os.Create("data")
	ioutil.WriteFile("data",data,0)

}


