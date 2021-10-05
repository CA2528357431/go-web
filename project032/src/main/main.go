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

func deletebyid(db *sql.DB,id int) {
	sql := "delete from post where id = $1"
	_,err := db.Exec(sql,id)
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

	deletebyid(db,9)
	
}
