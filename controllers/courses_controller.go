package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/models"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func SearchCourse(c *gin.Context) {

	query := strings.TrimSpace(c.Query("q"))
	results, err := services.SearchCourse(query)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CourseSearchResponse{
		Results: results,
	})
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	courseDto, err := services.GetCourse(courseID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func AddCourse(c *gin.Context) {
	var body dto.CourseDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Validar Datos en services? (Ver tema de validacion de InstructorID)

	courseDto, err := services.AddCourse(body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courseDto)

}

func UpdateOneCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var body dto.CourseDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course name cannot be empty"})
		return
	}

	if len(body.Description) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course description must be at least 10 characters long"})
		return
	}

	if len(body.Category) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course category cannot be empty"})
		return
	}

	if len(body.Requirements) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course requirements cannot be empty"})
		return
	}

	if body.Length <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course length must be a positive number of hours"})
		return
	}

	if !strings.HasPrefix(body.ImageURL, "http") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL"})
		return
	}

	datos := models.Course{}
	if err := db.GetDB().First(&datos, courseID); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No course found with that ID"})
		return
	} else {

		datos.Name = body.Name
		datos.Description = body.Description
		datos.Category = body.Category
		datos.Requirements = body.Requirements
		datos.Length = body.Length
		datos.ImageURL = body.ImageURL
		datos.LastUpdated = db.GetDB().NowFunc()

		db.GetDB().Save(&datos)
		c.JSON(http.StatusOK, gin.H{"Mensaje": "Course updated successfully"})
	}

}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	datos := models.Course{}
	if err := db.GetDB().First(&datos, courseID); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No course found with that ID"})
		return
	} else {
		db.GetDB().Delete(&datos)
		c.JSON(http.StatusOK, gin.H{"Mensaje": "Course deleted successfully"})
	}

}
