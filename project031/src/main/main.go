package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Content string
	Username string
	Id int
}

func (p *Post)changebyid(db *sql.DB,id int) {
	sql := "update post set username = $2, content = $3 where id = $1"
	_,err := db.Exec(sql,id,p.Username,p.Content)
	// 不关心结果，用exec
	if err!=nil{
		fmt.Println(err)
	}
}


func connect() *sql.DB {
	db,err := sql.Open("postgres","user=postgres dbname=post_test password=131128287 sslmode=disable")
	if err!=nil{
		fmt.Println(err)
	}
	return db
}
func main() {
	db := connect()

	p := Post{
		Content:  "looking forward to one",
		Username: "wyx",
	}

	p.changebyid(db,8)
	
}
