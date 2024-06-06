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

func SelectCoursesWithFilter(query string) ([]models.Course, error) {
	var courses []models.Course
	result := db.GetDB().Where("name LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}
