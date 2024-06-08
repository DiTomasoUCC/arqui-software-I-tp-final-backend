package controllers

import (
	"net/http"
	"regexp"
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
	c.JSON(200, gin.H{
		"mensaje": "metodo POST",
	})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request body"})
		return
	}

	// Comprehensive data validation

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

	// URL validation (consider using a dedicated library for robustness)
	if body.ImageURL != "" && !isValidURL(body.ImageURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image URL format"})
		return
	}

	datos := models.Course{}
	if err := db.GetDB().First(&datos, courseID); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No course found with that ID"})
		return
	} else {
		datos.Name = body.Name
		db.GetDB().Save(&datos)
		c.JSON(http.StatusOK, gin.H{"Mensaje": "Course updated successfully"})
	}
	// Update course data
	datos.Name = body.Name
	datos.Description = body.Description
	datos.Category = body.Category
	datos.Requirements = body.Requirements
	datos.Length = body.Length
	datos.ImageURL = body.ImageURL // Assuming ImageURL is a field in models.Course

	if err := db.GetDB().Save(&datos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while updating course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// isValidURL checks if the provided string is a valid URL using a basic regular expression
func isValidURL(url string) bool {
	// Regular expression to match basic URL format
	regex := `^((http|https)://)(www\.)?[a-zA-Z0-9@:%._\+~#?&//=]*$`

	// Create a Regex struct and compile the regex pattern
	r, err := regexp.Compile(regex)
	if err != nil {
		// Handle potential compilation errors
		return false
	}

	// Match the input string against the regex
	return r.MatchString(url)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"mensaje": "metodo DELETE / id=" + id,
	})
}
