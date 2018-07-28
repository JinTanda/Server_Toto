package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:0813@/Toto")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM classes")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next(){
		var id int
		var name sql.NullString
		var teacher sql.NullString
		var room sql.NullString
		var class_time sql.NullString
		var university sql.NullString
		var document sql.NullString
		if err := rows.Scan(&id, &name, &teacher, &room, &class_time, &university, &document); err != nil {
			panic(err.Error())
		}
		fmt.Println(id,name,teacher,room,class_time,university,document)
	}
}