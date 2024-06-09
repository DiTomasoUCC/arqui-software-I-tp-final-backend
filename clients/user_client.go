package clients

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func SelectUserByID(id int) (models.User, error) {
	var user models.User
	result := db.GetDB().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
