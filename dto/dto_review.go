package dto

type ReviewDto struct {
	CourseId int    `json:"course_id"`
	Comment  string `json:"comment"`
	Stars    int    `json:"stars"`
}

type ReviewsDto []ReviewDto
