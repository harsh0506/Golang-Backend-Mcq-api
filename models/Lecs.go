package models

import "gorm.io/gorm"

type LEC struct {
	gorm.Model
	College    string
	Department string
	Cource     string
	Year       string
	Sem        string
	Div        string
	Lec        string
	Start      string
	End        string
}
