package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
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
