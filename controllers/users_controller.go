package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/middleware"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	cook, err := c.Cookie("auth")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	valid := middleware.ValidateJWT(cook)

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	userID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	userDto, err := services.GetUser(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userDto)

}

func UserRegister(c *gin.Context) {
	var body dto.UserDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDto, err := services.CreateUser(body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userDto)

}

func UserLogin(c *gin.Context) {
	var body dto.LoginDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponseDto, err := services.LoginUser(body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("auth", loginResponseDto.Token, 3600*24, "", "", false, true) //3600 seconds = 1 hour

	c.JSON(http.StatusOK, loginResponseDto)

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = services.DeleteUser(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func GetUserCourses(c *gin.Context) {

}

func Logout(c *gin.Context) {
	c.SetCookie("auth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
}
