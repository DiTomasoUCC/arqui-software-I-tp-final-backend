package models

type Review struct {
	Id      int    `gorm:"primaryKey"`
	Comment string `gorm:"varchar(500);not null"`
	Stars   int    `gorm:"not null"`
}
