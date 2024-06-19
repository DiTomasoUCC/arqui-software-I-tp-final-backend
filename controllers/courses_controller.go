package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/middleware"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

func SearchCourse(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	category := strings.TrimSpace(c.Query("category"))

	results, err := services.SearchCourse(query, category)

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
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	courseDto, err := services.GetCourseWithBool(1, courseID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func AddCourse(c *gin.Context) {
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

	c.JSON(http.StatusCreated, courseDto)
}

func UpdateOneCourse(c *gin.Context) {
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

	courseDto, err := services.UpdateCourse(courseID, body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func DeleteCourse(c *gin.Context) {
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
	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	err = services.DeleteCourse(courseID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
