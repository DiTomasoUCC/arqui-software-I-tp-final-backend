package models

type Course struct {
	Id           int    `gorm:"primaryKey"`
	Name         string `gorm:"varchar(500);not null"`
	Description  string `gorm:"varchar(500);not null"`
	Instructor   int    `gorm:"varchar(500);not null"`
	Category     string `gorm:"varchar(500);not null"`
	Requirements string `gorm:"varchar(500);not null"`
	Lenght       int    `gorm:"not null"`

	Users   []User   `gorm:"many2many:CourseActivity"`
	Reviews []Review `gorm:"foreignKey:CourseID"`
}
