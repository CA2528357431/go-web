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

func getbyid(db *sql.DB,id int) Post  {
	sql := "select id, username, content from post where id = $1"
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}
	p := Post{}
	err = stmt.QueryRow(id).Scan(&p.Id,&p.Username,&p.Content)
	if err!=nil{
		fmt.Println(err)
	}
	return p

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
	p := getbyid(db,5)
	fmt.Println(p)

}
