package models

// import (
// 	"time"
// )
type User struct{
	ID string `json:"id"`
	University string `json:university"`
	Name string `json:"name"`
	Email string `json:"mail"`
	Password string `json:"password"`
	Department string `json:"department"`
	Major string `json:"major"`
}

type Users []User
