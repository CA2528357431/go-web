package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Acomment struct {

	Id int `xml:"id,attr"`

	Content string `xml:"content"`

	Username string `xml:"username"`
}

type Bcomment struct {

	Id int `xml:"id,attr"`

	Content string `xml:"content"`
}

type Post struct {

	XMLName xml.Name `xml:"mypost"`
	// 手动设置名称
	// 否则自动设为Struct的名字

	Id int `xml:"id,attr"`
	
	Content string `xml:"content"`

	Username string `xml:"username"`

	Acomments []Acomment `xml:"comments>acomment"`

	Bcomments []Bcomment `xml:"comments>bcomment"`

}

func main() {
	
	p := Post{
		Id:        7,
		Content:   "hello world",
		Username:  "caoan",
		Acomments: []Acomment{
			{
				Id:       1,
				Content:  "well",
				Username: "caoan",
			},
			{
				Id:       2,
				Content:  "honey",
				Username: "zyl",
			},
		},
		Bcomments: []Bcomment{
			{
				Id:      6,
				Content: "my lover",
			},
			{
				Id:      8,
				Content: "i like it",
			},
		},
	}

	data,err := xml.MarshalIndent(p,"","\t")
	// 更加美观的序列化
	// 第二个参数是主标签的前缀
	// 第三个参数是表示标签从属关系的符号

	if err != nil{
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("test.xml",data,0)

	if err != nil{
		fmt.Println(err)
		return
	}

}