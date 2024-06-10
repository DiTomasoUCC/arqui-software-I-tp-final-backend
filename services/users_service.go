package services

import (
	"fmt"
	"os"
	"time"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(email string, username string, id int) (string, error) {
	errorVariable := godotenv.Load()
	if errorVariable != nil {
		panic("Error loading .env file")
	}

	jwtKey := os.Getenv("SECRET_JWT")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"username": username,
		"id":       id,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(), //24 hours
	})

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, err
}

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

func LoginUser(loginDto dto.LoginDto) (dto.LoginResponseDto, error) {
	user, err := clients.SelectUserByEmail(loginDto.Email)

	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("error getting course from DB: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginDto.Password))

	if err != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("pass: "+loginDto.Password+"error getting course from DB: %w", err)
	}

	jwtKey, errJWT := GenerateJWT(user.Email, user.UserName, user.ID)

	if errJWT != nil {
		return dto.LoginResponseDto{}, fmt.Errorf("error getting course from DB: %w", errJWT)
	}

	return dto.LoginResponseDto{
		UserName: loginDto.Password,
		Token:    jwtKey,
	}, nil
}

func DeleteUser(id int) error {
	err := clients.DeleteUser(id)

	if err != nil {
		return fmt.Errorf("error getting course from DB: %w", err)
	}
	return nil
}
