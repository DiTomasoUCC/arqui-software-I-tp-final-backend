package models

import (
	"time"
)

// Review represents a review of a course.
type Review struct {
	ID           int `gorm:"primary_key"`
	CourseID     int // Foreign key referencing Course.ID
	UserID       int // Foreign key referencing User.ID
	Comment      string
	Stars        int
	CreationTime time.Time
	LastUpdated  time.Time
}
