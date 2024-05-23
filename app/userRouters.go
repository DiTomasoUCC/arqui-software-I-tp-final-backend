package routes

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.GetSingleUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateOneUser).Methods("PUT")
}
