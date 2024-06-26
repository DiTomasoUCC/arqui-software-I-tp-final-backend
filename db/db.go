package db

import (
	"fmt"
	"log"
	"os"
	"time"

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
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_SERVER") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("DSN: %s\n", dsn) // Print DSN for debugging

	// Retry logic for database connection
	var retryCount int
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if retryCount >= 5 {
			return fmt.Errorf("failed to connect database after retries: %w", err)
		}
		retryCount++
		log.Printf("Database connection failed. Retrying... (%d/5)\n", retryCount)
		time.Sleep(5 * time.Second)
	}

	// Log the database connection success
	fmt.Println("Database connection successful")

	// Perform auto-migration
	err = db.AutoMigrate(&models.User{}, &models.Course{}, &models.Subscription{}, &models.Activity{}, &models.Comment{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	// Log the auto-migration success
	fmt.Println("Auto-migration successful")
	return nil
}
