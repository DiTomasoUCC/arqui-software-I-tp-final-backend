package routes

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/Users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/User", controllers.AddUser).Methods("POST")
	router.HandleFunc("/User/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/User/{id}", controllers.GetSingleUser).Methods("GET")
	router.HandleFunc("/User/{id}", controllers.UpdateOneUser).Methods("PUT")
}
