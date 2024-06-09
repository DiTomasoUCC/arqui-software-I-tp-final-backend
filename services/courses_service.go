package services

import (
	"fmt"
	"strings"

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

func SearchCourse(query string) ([]dto.CourseDto, error) {
	trimmed := strings.TrimSpace(query)

	courses, err := clients.SelectCoursesWithFilter(trimmed)

	if err != nil {
		return nil, fmt.Errorf("error getting course from DB: %w", err)
	}

	results := make([]dto.CourseDto, 0)

	for _, course := range courses {
		results = append(results, dto.CourseDto{
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
		})
	}

	return results, nil
}

func AddCourse(courseDto dto.CourseDto) (dto.CourseDto, error) {
	course, err := clients.CreateCourse(courseDto.Name, courseDto.Description, courseDto.InstructorId, courseDto.Category, courseDto.Requirements, courseDto.Length, courseDto.ImageURL)

	if err != nil {
		return dto.CourseDto{}, fmt.Errorf("error creating course in DB: %w", err)
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
