package models

import "time"

type Subscription struct {
	ID           int `gorm:"primary_key"`
	CourseID     int //Foreign key referencing Course ID
	UserID       int //Foreign key referencing Course ID
	CreationTime time.Time
}
