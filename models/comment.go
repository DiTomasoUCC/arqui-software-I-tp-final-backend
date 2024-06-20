package models

import (
	"time"
)

// Comment represents a comment of a course.
type Comment struct {
	ID           int `gorm:"primary_key"`
	CourseID     int // Foreign key referencing Course.ID
	UserID       int // Foreign key referencing User.ID
	Comment      string
	CreationTime time.Time
}
