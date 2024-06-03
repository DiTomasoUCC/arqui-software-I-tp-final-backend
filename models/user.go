package models

type User struct {
	Id       int      `gorm:"primaryKey"`
	UserName string   `gorm:"varchar(600);not null"`
	Name     string   `gorm:"varchar(500);not null"`
	LastName string   `gorm:"varchar(500);not null"`
	UserType bool     `gorm:"not null"`
	Reviews  []Review `gorm:"foreignKey:UserID"`
	Courses  []Course `gorm:"many2many:CourseActivity"`
}
