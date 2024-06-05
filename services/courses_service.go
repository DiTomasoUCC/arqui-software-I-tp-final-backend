package services

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

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
