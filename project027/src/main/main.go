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



	posts := []post{
		{"hello",1,"caoan"},
		{"i'm neo here",2,"caoan"},
		{"i love you",3,"wangyuxiao"},
	}
	filenames := []string{
		"post1","post2","post3",
	}

	for i:=0;i<3;i++{
		buffer := &bytes.Buffer{}
		encoder := gob.NewEncoder(buffer)
		encoder.Encode(posts[i])
		os.WriteFile(filenames[i],buffer.Bytes(),1)
		// ioutil 也有一样的函数
	}

	fmt.Println("end")
}