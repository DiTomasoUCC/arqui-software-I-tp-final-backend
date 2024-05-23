package main

import (
	"log"
	"net/http"

	routes "github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/app"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page!\n"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", YourHandler)
	routes.CourseRoutes(router)
	routes.UserRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
