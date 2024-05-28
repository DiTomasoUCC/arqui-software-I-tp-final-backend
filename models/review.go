package models

type Review struct {
	Id      int    `gorm:"primaryKey"`
	Comment string `gorm:"varchar(500);not null"`
	Stars   int    `gorm:"not null"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	CourseID uint
	Course   Course `gorm:"foreignKey:CourseID"`
}
