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

// UserToCln user info send to cln structure
type UserToCln struct {
	Name      string `json:"name" gorm:"varchar(20);not null"`
	Phone     string `json:"phone" gorm:"varchar(20);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}

func ToUserDto(u User) UserToCln {
	return UserToCln{
		Name:      u.Name,
		Phone:     u.Phone,
		CreatedAt: Time(u.CreatedAt),
		UpdatedAt: Time(u.UpdatedAt),
	}
}
