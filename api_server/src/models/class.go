package models

type Class struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Teacher string `json:"teacher"`
	Room string `json:"room"`
	Class_time string `json:"time"`
	University string `json:"university"`
	Document string `json:"document"`
}

type Classes []Class