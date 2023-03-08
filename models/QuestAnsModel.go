package models

import "gorm.io/gorm"

type MCQ struct {
	gorm.Model
	Question string
	O1       string
	O2       string
	O3       string
	O4       string
	Ans      string
	Sub      string
}
