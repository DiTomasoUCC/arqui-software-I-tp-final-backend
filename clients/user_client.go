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

func CreateUser(email string, username string, firstName string, lastName string, userType bool, passwordHash string) (models.User, error) {
	user := models.User{
		Email:        email,
		UserName:     username,
		FirstName:    firstName,
		LastName:     lastName,
		UserType:     userType,
		PasswordHash: passwordHash,
		CreationTime: db.GetDB().NowFunc(),
		LastUpdated:  db.GetDB().NowFunc(),
	}
	result := db.GetDB().Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func DeleteUser(id int) error {
	result := db.GetDB().Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
