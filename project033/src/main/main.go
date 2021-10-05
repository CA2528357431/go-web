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

func gets(db *sql.DB,name string) []Post  {
	sql := "SELECT id, username, content FROM post WHERE username = $1 "
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}
	posts := []Post{}
	rows,err := stmt.Query(name)
	if err!=nil{
		fmt.Println(err)
	}
	// rows是迭代器
	// next进行迭代

	for rows.Next(){
		p := Post{}
		err = rows.Scan(&p.Id,&p.Username,&p.Content)
		if err!=nil{
			fmt.Println(err)
		}
		posts = append(posts,p)
	}

	rows.Close()

	return posts
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
	posts := gets(db,"wyx")
	for i,p := range posts{
		fmt.Println(i,": ",p)
	}

}
