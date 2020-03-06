package models

import "github.com/jinzhu/gorm"

type User struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Role 		uint8	`json:"role"`
	gorm.Model
}

