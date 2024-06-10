package dto

import "time"

type UserDto struct {
	UserId       int    `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	FirstName    string `json:"name"`
	LastName     string `json:"last_name"`
	UserType     bool   `json:"usertype"`
	PasswordHash string `json:"password_hash"`
	CreationTime time.Time
	LastUpdated  time.Time
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	UserId int    `json:"id"`
	Token  string `json:"token"`
}
