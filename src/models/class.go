package models

type Class struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Teacher string `json:"teacher"`
	Room string `json:"room"`
	Time string `json:"time"`
	University string `json:"university"`
	Document string `json:"document"`
}

type Classes []Class