package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Content string
	Username string
	id int
}

func (p *Post)createcomment(db *sql.DB)  {
	sql := "insert into comment (username, content, post_id) values ($1, $2, $3)"
	_,err := db.Exec(sql,p.Username,p.Content,p.id)
	if err!=nil{
		fmt.Println(err)
	}
}

func (p *Post)createpost(db *sql.DB)  {
	sql := "insert into post (username, content) values ($1, $2)"
	_,err := db.Exec(sql,p.Username,p.Content)
	if err!=nil{
		fmt.Println(err)
	}
}

func deletepost(db *sql.DB,id int) {
	sql := "delete from post where id = $1"
	_,err := db.Exec(sql,id)
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
		Content:  "hello,people",
		Username: "zzy",
		id: 	  9,
	}
	p.createpost(db)
	p.createcomment(db)
	deletepost(db,9)


}
