package dto

import "time"

type CommentDto struct {
	ID           	int    			`json:"comment_id"`
	CourseId 			int    			`json:"course_id"`
	UserId   			int    			`json:"user_id"`
	Comment  			string 			`json:"comment"`
	CreationTime 	time.Time 	`json:"creation_time"`
}

type CommentsDto []CommentDto