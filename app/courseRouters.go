package routes

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"
	"github.com/gorilla/mux"
)

func CourseRoutes(router *mux.Router) {
	router.HandleFunc("/courses", controllers.GetAllCourses).Methods("GET")
	router.HandleFunc("/course", controllers.AddCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controllers.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/course/{id}", controllers.GetSingleCourse).Methods("GET")
	router.HandleFunc("/course/{id}", controllers.UpdateOneCourse).Methods("PUT")
}
