package clients

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func CreateSubscription(courseId int, userId int) (models.Subscription, error) {
	subscription := models.Subscription{
		CourseID:     courseId,
		UserID:       userId,
		CreationTime: db.GetDB().NowFunc(),
	}
	result := db.GetDB().Create(&subscription)
	if result.Error != nil {
		return models.Subscription{}, result.Error
	}
	return subscription, nil
}

func GetSubscribedUsers(courseId int) ([]models.User, error) {
	var users []models.User
	result := db.GetDB().Table("users").Select("users.*").Joins("JOIN subscriptions ON users.id = subscriptions.user_id").Where("subscriptions.course_id = ?", courseId).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func SelectSubscription(courseId int, userId int) models.Subscription {
	var subscription models.Subscription
	result := db.GetDB().Where("course_id = ? AND user_id = ?", courseId, userId).First(&subscription)
	if result.Error != nil {
		return models.Subscription{}
	}
	return subscription
}
