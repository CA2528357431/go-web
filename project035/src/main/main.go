package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Content string
	Username string
	Id int
}

func (p *Post)create(db *sqlx.DB)  {
	sql := "insert into post (username, content) values ($1,$2)"
	err := db.QueryRowx(sql,p.Username,p.Content)
	if err!=nil{
		fmt.Println(err)
	}
}

func get(db *sqlx.DB,id int) Post  {
	sql := "select id, username, content from post where id = $1"
	p := Post{}
	row := db.QueryRowx(sql,id)
	err := row.StructScan(&p)

	// struct的匹配是建立在 数据库字段 与 struct tag/field之间的

	if err!=nil{
		fmt.Println(err)
	}

	return p
}

func gets(db *sqlx.DB,name string) []Post  {
	sql := "SELECT id, username, content FROM post WHERE username = $1 "
	rows,err := db.Queryx(sql,name)
	if err!=nil{
		fmt.Println(err)
	}
	posts := []Post{}

	for rows.Next(){
		p := Post{}
		err = rows.StructScan(&p)
		if err!=nil{
			fmt.Println(err)
		}
		posts = append(posts,p)
	}

	rows.Close()

	return posts
}

func (p *Post)change(db *sqlx.DB,id int) {
	sql := "update post set username = $2, content = $3 where id = $1"
	_ = db.QueryRowx(sql,id,p.Username,p.Content)
}

func remove(db *sqlx.DB,id int) {
	sql := "delete from post where id = $1"
	err := db.QueryRowx(sql,id)
	if err!=nil{
		fmt.Println(err)
	}
}

func connect() *sqlx.DB {
	db,err := sqlx.Open("postgres","user=postgres dbname=post_test password=131128287 sslmode=disable")
	if err!=nil{
		fmt.Println(err)
	}
	return db
}
func main() {
	//	db := connect()

	//	p := Post{
	//		Content:  "hello,anan",
	//		Username: "wyx",
	//	}
	//	p.create(db)

	//	p := get(db,9)
	//	fmt.Println(p)

	// 	posts := gets(db,"wyx")
	//	fmt.Println(posts)

	// 	remove(db,10)

	//	p := Post{
	//		Content:  "anan,miss you",
	//		Username: "wyx",
	//	}
	//	p.change(db,9)

}

// 此外还有gorm等库操作数据库
