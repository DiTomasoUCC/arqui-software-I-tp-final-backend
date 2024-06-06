package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func GetAllCourses(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "metodo GET ALL courses",
	})
}

func GetSingleCourse(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo GET / id=" + id,
	})

}

func AddCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "metodo POST",
	})
}

func UpdateOneCourse(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo PUT / id=" + id,
	})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo DELETE / id=" + id,
	})
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := services.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	courseDto := dto.ToCourseDto(course) // Convert course model to DTO
	c.JSON(http.StatusOK, courseDto)
}
