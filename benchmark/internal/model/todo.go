package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model    `swaggerignore:"true"`
	Title         string `json:"title"`
	Description   string `json:"description"` 
	Completed     bool   `json:"completed"`
	UserID        uint   `json:"user_id"`
	User          User   `json:"-" gorm:"foreignKey:UserID"` 
}