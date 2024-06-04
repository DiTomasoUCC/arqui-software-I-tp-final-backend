package controllers

import (
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
		"mensaje": "metodo GET single course / id= " + id,
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
