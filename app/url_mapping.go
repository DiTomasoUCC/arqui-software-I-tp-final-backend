package app

import "github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"

var prefix = "/api/v1/"

func mapUrls() {
	//Courses endpoints
	router.GET(prefix+"courses/search", controllers.SearchCourse)
	router.GET(prefix+"course/:id", controllers.GetCourse)
	router.POST(prefix+"course", controllers.AddCourse)
	router.PUT(prefix+"course/:id", controllers.UpdateOneCourse)
	router.DELETE(prefix+"course/:id", controllers.DeleteCourse)

	//Users endpoints
	router.GET(prefix+"user/:id", controllers.GetUser)
	router.DELETE(prefix+"user/:id", controllers.DeleteUser)
	router.POST(prefix+"user/register", controllers.UserRegister)
	router.POST(prefix+"user/login", controllers.UserLogin)
	router.GET(prefix+"user/logout", controllers.Logout)
	router.GET(prefix+"user/courses/:user_id", controllers.GetUserCourses)

	//Subscriptions endpoints
	router.GET(prefix+"subscriptions/:course_id", controllers.GetSubscribedUsers)
	router.POST(prefix+"subscription", controllers.AddSubscription)

	// Comments endpoints
	router.GET(prefix+"comments/:course_id", controllers.GetComments)
	router.POST(prefix+"comments", controllers.AddComment)

	//Upload example
	router.POST(prefix+"upload", controllers.UploadFile)
}
