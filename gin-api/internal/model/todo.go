package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model 		`swaggerignore:"true"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
