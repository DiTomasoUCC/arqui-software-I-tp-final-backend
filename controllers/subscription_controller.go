package controllers

import (
	"net/http"
	"strconv"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/dto"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/services"
	"github.com/gin-gonic/gin"
)

// func GetCourseSubscriptions(c *gin.Context) {
// 	id := c.Param("course_id")

// }

func GetSubscribedUsers(c *gin.Context) {
	id := c.Param("course_id")

	courseID, err := strconv.Atoi(id) // Convert string ID to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	subscribedUsers, err := services.GetSubscribedUsers(courseID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.CourseSubscriptionsResponse{
		Results: subscribedUsers,
	})

}

func AddSubscription(c *gin.Context) {
	var body dto.SubscriptionDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionDto, err := services.AddSubscription(body)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, subscriptionDto)

}
