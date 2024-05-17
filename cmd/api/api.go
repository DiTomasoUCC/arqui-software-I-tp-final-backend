package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/service/user"
	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, nil)
}
