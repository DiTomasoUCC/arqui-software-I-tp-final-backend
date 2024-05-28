package models

type Activity struct {
	Id                  int    `gorm:"primaryKey"`
	ActivityDescription string `gorm:"varchar(500);not null"`
	UserID              uint   // Foreign key referencing User table
	CourseID            uint   // Foreign key referencing Course table
}
