package clients

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

func GetComments(courseID int) ([]models.Comment, error) {
	var comments []models.Comment
	result := db.GetDB().Where("course_id = ?", courseID).Find(&comments)
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