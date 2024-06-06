package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

func GetCourse(id int) (dto.CourseDto, error) {
	course, err := clients.SelectCourseByID(id)

	if err != nil {
		return dto.CourseDto{}, fmt.Errorf("error getting course from DB: %w", err)
	}
	return dto.CourseDto{
		Id:           course.ID,
		Name:         course.Name,
		Description:  course.Description,
		InstructorId: course.InstructorID,
		Category:     course.Category,
		Requirements: course.Requirements,
		Length:       course.Length,
		ImageURL:     course.ImageURL,
		CreationTime: course.CreationTime,
		LastUpdated:  course.LastUpdated,
	}, nil
}
