package validator

import "demo/model"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDto(u model.User) UserDto {
	return UserDto{
		Name:  u.Name,
		Phone: u.Phone,
	}
}
