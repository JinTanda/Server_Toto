package models

// import (
// 	"time"
// )
type User struct{
	ID string `json:"id"`
	University string `json:university"`
	Name string `json:"name"`
	Email string `json:"mail"`
	password string `json:"password"`
	department string `json:"department"`
	major string `json:"major"`
}

type Users []User
