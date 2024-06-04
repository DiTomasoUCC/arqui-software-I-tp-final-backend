package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	mapUrls()

	logrus.Info("Starting server")
	router.Run(":8080")
}
