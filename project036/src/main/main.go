package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {

	Id int `xml:"id,attr"`
	// attr表示数据在属性里
	
	Content string `xml:"content"`

	Username string `xml:"username"`

	XMLName xml.Name `xml:"mypost"`
	// xml主名称
	// 名字必须是XMLName

	XMLori string `xml:",innerxml"`
	// `xml:",innerxml"`获取原文

	Data []byte `xml:",chardata"`
	// 各行无标签数据之和
	

}

func main() {
	data,err := ioutil.ReadFile("test.xml")

	if err != nil{
		fmt.Println(err)
		return
	}

	p := Post{}

	err = xml.Unmarshal(data,&p)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(p)

}