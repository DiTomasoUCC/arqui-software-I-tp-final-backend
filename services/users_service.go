package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

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
