package clients

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func SelectCourseByID(id int64) (models.Course, error) {
	var course models.Course
	result := db.First(&course, id)

	if result.Error != nil {
		return models.Course{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return course, nil
}
