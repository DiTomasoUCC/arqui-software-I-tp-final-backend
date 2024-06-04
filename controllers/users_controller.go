package controllers

import (
	"net/http"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "metodo GET ALL Users",
	})
}

func GetSingleUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo GET single User / id= " + id,
	})
}

func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "metodo POST",
	})
}

func UpdateOneUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo PUT / id=" + id,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo DELETE / id=" + id,
	})
}

// -----------------------------------------------------------------------------------------
type UserController struct {
	userService services.UserService
}

// NewUserController creates a new user controller.
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterUser handles user registration requests.
func (uc *UserController) RegisterUser(c *gin.Context) {
	// Extract user data from request body
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call user service to register the user
	if err := uc.userService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login requests.
func (uc *UserController) LoginUser(c *gin.Context) {
	// Extract username and password from request body
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call user service to validate credentials and get JWT token
	token, err := uc.userService.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Send the JWT token in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
