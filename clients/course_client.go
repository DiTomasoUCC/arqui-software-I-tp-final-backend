package clients

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func SelectCourseByID(id int) (models.Course, error) {
	var course models.Course
	result := db.GetDB().Where("id = ?", id).First(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}
