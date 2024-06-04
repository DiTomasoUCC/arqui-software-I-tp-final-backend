package app

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"
)

func mapUrls() {
	// User routes
	userRouter := router.Group("/api/v1/users")
	userRouter.POST("/register", controllers.RegisterUser)
	userRouter.POST("/login", controllers.LoginUser)

	// Other routes... (replace with your existing mappings)
}
