package app

import "github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"

var prefix = "/api/v1/"

func mapUrls() {

	//Courses endpoints
	router.GET(prefix+"courses/search", controllers.SearchCourse) //example: http://localhost:8080/api/v1/courses/search?q=Go_lang_Course
	router.GET(prefix+"course/:id", controllers.GetCourse)
	router.POST(prefix+"course", controllers.AddCourse)
	router.PUT(prefix+"course/:id", controllers.UpdateOneCourse)
	router.DELETE(prefix+"course/:id", controllers.DeleteCourse)

	//Users endpoints
	router.GET(prefix+"user/:id", controllers.GetUser)
	router.POST(prefix+"user/register", controllers.UserRegister)
	router.POST(prefix+"user/login", controllers.UserLogin)
	router.PUT(prefix+"user/:id", controllers.UpdateUser)
	router.DELETE(prefix+"user/:id", controllers.DeleteUser)
}
