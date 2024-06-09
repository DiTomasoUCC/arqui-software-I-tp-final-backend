package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	cost := 10
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes)
}

func GetUser(id int) (dto.UserDto, error) {
	user, err := clients.SelectUserByID(id)

	if err != nil {
		return dto.UserDto{}, fmt.Errorf("error getting course from DB: %w", err)
	}
	return dto.UserDto{
		UserId:       user.ID,
		Email:        user.Email,
		Username:     user.UserName,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserType:     user.UserType,
		PasswordHash: user.PasswordHash,
		CreationTime: user.CreationTime,
		LastUpdated:  user.LastUpdated,
	}, nil
}

func CreateUser(userDto dto.UserDto) (dto.UserDto, error) {

	userDto.PasswordHash = HashPassword(userDto.PasswordHash)

	user, err := clients.CreateUser(userDto.Email, userDto.Username, userDto.FirstName, userDto.LastName, userDto.UserType, userDto.PasswordHash)

	if err != nil {
		return dto.UserDto{}, fmt.Errorf("error getting course from DB: %w", err)
	}
	return dto.UserDto{
		UserId:       user.ID,
		Email:        user.Email,
		Username:     user.UserName,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		UserType:     user.UserType,
		PasswordHash: user.PasswordHash,
		CreationTime: user.CreationTime,
		LastUpdated:  user.LastUpdated,
	}, nil
}

func DeleteUser(id int) error {
	err := clients.DeleteUser(id)

	if err != nil {
		return fmt.Errorf("error getting course from DB: %w", err)
	}
	return nil
}
