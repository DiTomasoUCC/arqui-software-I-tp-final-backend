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

func CreateCourse(name string, desc string, instruct int, category string, req string, length int, img string) (models.Course, error) {
	course := models.Course{
		Name:         name,
		Description:  desc,
		InstructorID: instruct,
		Category:     category,
		Requirements: req,
		Length:       length,
		ImageURL:     img,
		CreationTime: db.GetDB().NowFunc(),
		LastUpdated:  db.GetDB().NowFunc(),
	}
	result := db.GetDB().Create(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func UpdateCourse(id int, name string, desc string, category string, req string, length int, img string) (models.Course, error) {
	course := models.Course{
		Name:         name,
		Description:  desc,
		Category:     category,
		Requirements: req,
		Length:       length,
		ImageURL:     img,
		LastUpdated:  db.GetDB().NowFunc(),
	}
	result := db.GetDB().Model(&models.Course{}).Where("id = ?", id).Updates(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}
