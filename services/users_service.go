package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
)

// UserService provides methods for user management.
type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (string, error)
}

type userService struct {
	db *gorm.DB
}

// NewUserService creates a new user service.
func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// RegisterUser registers a new user in the database.
func (s *userService) RegisterUser(user *models.User) error {
	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)

	// Create a new user record in the database
	return s.db.Create(user).Error
}

// LoginUser validates the user credentials and returns a JWT token if successful.
func (s *userService) LoginUser(username, password string) (string, error) {
	// Find the user by username
	var user models.User
	if err := s.db.First(&user, "user_name = ?", username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate a JWT token for the user
	token, err := generateJWT(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateJWT generates a JSON Web Token (JWT) for the user.
func generateJWT(userID uint) (string, error) {
	// Replace with your own secret key
	secretKey := []byte("your_secret_key")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "your_application_name",               // Issuer
		"sub": userID,                                // Subject
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires after 24 hours
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
