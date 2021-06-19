package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"varchar(20);not null"`
	Phone    string `json:"phone" gorm:"varchar(20);not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
}
