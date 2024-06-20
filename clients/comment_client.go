package clients

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func GetComments(courseID int) ([]models.CommentQuery, error) {
	var comments []models.CommentQuery

	// Use Raw SQL query to join comments with users and select user_name
	result := db.GetDB().Raw(`
		SELECT comments.*, users.user_name AS user_name
		FROM comments
		LEFT JOIN users ON users.id = comments.user_id
		WHERE comments.course_id = ?
		ORDER BY comments.creation_time DESC
	`, courseID).Scan(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func CreateComment(courseID int, userID int, comment string) (models.Comment, error) {
	newComment := models.Comment{
		CourseID:     courseID,
		UserID:       userID,
		Comment:      comment,
		CreationTime: db.GetDB().NowFunc(),
	}
	result := db.GetDB().Create(&newComment)
	if result.Error != nil {
		return models.Comment{}, result.Error
	}
	return newComment, nil
}