package models

type Activity struct {
	Id                  int    `gorm:"primaryKey"`
	ActivityDescription string `gorm:"varchar(500);not null"`
}
