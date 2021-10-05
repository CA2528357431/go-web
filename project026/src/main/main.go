package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type post struct {
	Content string
	Id int
	User string
}

func main() {

	file,_ := os.Open("data.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	// 限制读取字段数
	// 如果为-1 不限制
	// 如果为0 按照第一条限制
	// 如果为正则读取该数的字段，若字段数不足，则报错

	records,_ := reader.ReadAll()
	posts := make([]post,0)

	for _,record := range records{
		id,_ := strconv.Atoi(record[0])
		p := post{
			Content: record[2],
			Id:      id,
			User:    record[1],
		}
		posts = append(posts,p)
	}
	for _,p := range posts{
		fmt.Println(p.Id)
		fmt.Println(p.User)
		fmt.Println(p.Content)
		fmt.Println("-------")
	}

}