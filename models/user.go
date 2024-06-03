package models

import (
	"time"
)

// User represents a user in the system.
type User struct {
	ID           int    `gorm:"primary_key"`
	UserName     string `gorm:"unique"`
	LastName     string
	UserType     bool // true for instructor, false for student
	PasswordHash string
	CreationTime time.Time
	LastUpdated  time.Time
}
