package services

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func GetCourseByID(id int) (models.Course, error) {
	var course models.Course
	result := db.GetDB().Where("id = ?", id).First(&course)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func Get(id int64) (dto.CourseDto, error) {
	course, err := clients.SelectCourseByID(id)
	if err != nil {
		return dto.CourseDto{}, err
	}

	return dto.CourseDto{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
	}, nil
}
