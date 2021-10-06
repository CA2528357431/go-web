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

	Id int `xml:"id,attr"`
	
	Content string `xml:"content"`

	Username string `xml:"username"`

	Acomments []Acomment `xml:"comments>acomment"`
	Bcomments []Bcomment `xml:"comments>bcomment"`
	// 树状嵌套


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