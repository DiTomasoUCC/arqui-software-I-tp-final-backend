package dto

import "time"

type CourseDto struct {
	Id           int    `json:"course_id"`
	Name         string `json:"course_name"`
	Description  string `json:"description"`
	InstructorId int    `json:"instructor_id"`
	Category     string `json:"category"`
	Requirements string `json:"requirements"`
	Length       int    `json:"length"`
	ImageURL     string
	CreationTime time.Time
	LastUpdated  time.Time
}

type CoursesDto []CourseDto
