// json反序列化为数据

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	data,err := ioutil.ReadFile("data")
	if err != nil{
		fmt.Println(err)
		return
	}
	var obj per

	err = json.Unmarshal(data,&obj)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(obj)

}

type per struct {
	Name string
	Age int
	School []string
}
