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

func (p *Post)create(db *sql.DB)  {
	sql := "insert into post_test (username, content) values ($1,$2)"
	_,err := db.Exec(sql,p.Username,p.Content)
	if err!=nil{
		fmt.Println(err)
	}
	// 方便,对结果不感兴趣的时候使用
}

func (p *Post)createNEO(db *sql.DB)  {
	sql := "insert into post (username, content) values ($1,$2) returning id"
	stmt,err := db.Prepare(sql)
	if err!=nil{
		fmt.Println(err)
	}

	err = stmt.QueryRow(p.Username,p.Content).Scan(&(p.id))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(p)
	// 方便,对结果感兴趣的时候使用
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
	fmt.Println("connect")
	p := Post{
		Content:  "hello,anan",
		Username: "wyx",
	}
	p.create(db)




	pp := Post{
		Content:  "nice to meet you,wyx",
		Username: "CA",
	}
	pp.createNEO(db)



}
