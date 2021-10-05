package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type post struct {
	Content string
	Id int
	User string
}

func main() {

	file,_ := os.Create("data.csv")
	defer file.Close()

	posts := []post{
		{
			Content: "hello",
			Id:      1,
			User:    "Cao An",
		},
		{
			Content: "nice to meet",
			Id:      2,
			User:    "caoan",
		},
		{
			Content: "i love you",
			Id:      3,
			User:    "Wang Yuxiao",
		},
	}

	writer := csv.NewWriter(file)
	for _,p := range posts{
		line := []string{strconv.Itoa(p.Id),p.User,p.Content}
		writer.Write(line)
	}
	writer.Flush()
	// 更新数据

}