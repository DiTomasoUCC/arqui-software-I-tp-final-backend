package dto

type CourseDto struct {
	CourseId     int     `json:"course_id"`
	CourseName   string  `json:"course_name"`
	Category     string  `json:"category"`
	Length       float32 `json:"length"`
	InstructorId int     `json:"instructor_id"`
	Requirements string  `json:"requirements"`
	Description  string  `json:"description"`
}

type CoursesDto []CourseDto
