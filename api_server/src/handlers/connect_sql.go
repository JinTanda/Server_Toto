package main

import (
	"database/sql"
	"log"
	"../models"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	var class models.Class
	class = GetClassByID(1)
	log.Println(class)

	var user models.User
	user.University = "NIT"
	user.Name = "よしの"
	user.Email = "yoshino.kai@itolab.nitech.ac.jp"
	user.Password = "hogehoge"
	user.Department = "工学部"
	user.Major = "Computer Science"
	PostUser(user)
}

func PostUser(user models.User){
	db,err := sql.Open("mysql","root:0813@/Toto")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	query := "SELECT count(email) FROM users WHERE EXISTS (SELECT * FROM users WHERE users.email = '" + user.Email + "')"
	
	var email string
	if err := db.QueryRow(query).Scan(&email); err != nil{
		panic(err.Error())
	}

	if email == "0"{
		query := "INSERT INTO users (university,name,email,password,department,major) VALUES ('" + user.University + "','" + user.Name + "','" + user.Email + "','" + user.Password + "','" + user.Department + "','" + user.Major + "')"
		row, err := db.Query(query)
		defer row.Close()
		if err != nil{
			panic(err.Error())
		}
		log.Println(query)
		log.Println("Insert successed")
	} else{
		log.Println("Insert failed. Email already exists")
	}



}


func GetClassByID(id int) models.Class{
	db,err := sql.Open("mysql","root:0813@/Toto")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM classes where id=?",id)
	defer row.Close()
	if err != nil {
		panic(err.Error())
	}

	row.Next()
	var class_id int
	var name string
	var teacher sql.NullString
	var room sql.NullString
	var class_time sql.NullString
	var university sql.NullString
	var document sql.NullString
	if err := row.Scan(&class_id, &name, &teacher, &room, &class_time, &university, &document); err != nil{
		panic(err.Error())
	}

	var class models.Class
	class.ID = class_id
	class.Name = name
	if teacher.Valid == true {class.Teacher=teacher.String} else{class.Teacher=""}
	if room.Valid == true {class.Room=room.String} else{class.Room=""}
	if class_time.Valid == true {class.Time=class_time.String} else{class.Time=""}
	if university.Valid == true {class.University=university.String} else{class.University=""}
	if document.Valid == true {class.Document=document.String} else{class.Document=""}
	log.Println("取り出した授業情報：?",class)
	return class
}