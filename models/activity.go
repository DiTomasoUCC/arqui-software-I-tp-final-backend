package models

import (
	"time"
)

// Activity represents an activity in a course.
type Activity struct {
	ID          int `gorm:"primary_key"`
	CourseID    int // Foreign key referencing Course.ID
	UserID      int // Foreign key referencing User.ID
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
