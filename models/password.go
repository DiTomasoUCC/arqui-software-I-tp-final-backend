package models

type Password struct {
	Id       int `gorm:"primaryKey"`
	UserID   uint
	Password string `gorm:"type:varchar(255);not null"`

	User User `gorm:"foreignKey:UserID"`
}
