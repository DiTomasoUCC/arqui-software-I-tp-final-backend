package services

import (
	"fmt"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/clients"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
)

func GetComments(courseID int) ([]dto.CommentDto, error) {
	comments, err := clients.GetComments(courseID)

	if err != nil {
		return nil, fmt.Errorf("error getting comments from DB: %w", err)
	}

	result := make([]dto.CommentDto, 0)

	for _, comment := range comments {
		result = append(result, dto.CommentDto{
			ID:           comment.ID,
			CourseId:     comment.CourseID,
			UserId:       comment.UserID,
			Comment:      comment.Comment,
			CreationTime: comment.CreationTime,
		})
	}

	return result, nil
}

func AddComment(body dto.CommentDto, user int) (dto.CommentDto, error) {
	comment, err := clients.CreateComment(body.CourseId, user, body.Comment)

	if err != nil {
		return dto.CommentDto{}, fmt.Errorf("error creating comment in DB: %w", err)
	}

	return dto.CommentDto{
		ID:           comment.ID,
		CourseId:     comment.CourseID,
		UserId:       comment.UserID,
		Comment:      comment.Comment,
		CreationTime: comment.CreationTime,
	}, nil
}