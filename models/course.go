package models

import (
	"time"
)

// Course represents a course in the system.
type Course struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Description  string
	InstructorID int // Foreign key referencing User.ID
	Category     string
	Requirements string
	Length       int // Duration in hours
	ImageURL     string
	CreationTime time.Time
	LastUpdated  time.Time
}

type CourseQuery struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Description  string
	InstructorID int // Foreign key referencing User.ID
	InstructorName string
	Category     string
	Requirements string
	Length       int // Duration in hours
	ImageURL     string
	CreationTime time.Time
	LastUpdated  time.Time
}
