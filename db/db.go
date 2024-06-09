package db

import (
	"fmt"
	"log"
	"os"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		err := ConnectDatabase()
		if err != nil {
			log.Panicf("failed to connect database: %v", err)
		}
	}
	return db
}

func ConnectDatabase() error {

	errorvariable := godotenv.Load()
	if errorvariable != nil {
		log.Fatalf("err loading: %v", errorvariable)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	fmt.Println(db)
	err = db.AutoMigrate(&models.User{}, &models.Course{}, &models.Subscription{}, &models.Activity{}, &models.Review{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}
