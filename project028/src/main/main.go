package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type post struct {
	Content string
	Id int
	User string
}

func main() {

	filenames := []string{
		"post1","post2","post3",
	}

	for i:=0;i<3;i++{
		data,_ := os.ReadFile(filenames[i])
		// ioutil 也有一样的函数
		buffer := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buffer)
		p := &post{}
		decoder.Decode(p)
		fmt.Println(*p)

	}

	fmt.Println("end")
}