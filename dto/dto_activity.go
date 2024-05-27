package dto

type ActivityDto struct {
	UserId      int    `json:"user_id"`
	CourseId    int    `json:"course_id"`
	Description string `json:"description"`
}

type ActivitiesDto []ActivityDto
