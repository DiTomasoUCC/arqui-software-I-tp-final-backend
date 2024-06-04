package dto

type UserDto struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserType bool   `json:"usertype"`
}

type UsersDto []UserDto