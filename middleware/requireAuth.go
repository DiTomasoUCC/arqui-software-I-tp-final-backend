package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ValidateJWT(header string) bool {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		return false

	}

	myKey := []byte(os.Getenv("SECRET_JWT"))

	if len(header) == 0 {
		return false
	}

	spliToken := strings.Split(header, ".")
	if len(spliToken) != 3 {
		return false
	}

	tk := strings.TrimSpace(header)

	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return myKey, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User

		db.GetDB().First(&user, claims["id"])

		return user.ID != 0

	} else {
		return false
	}
}
