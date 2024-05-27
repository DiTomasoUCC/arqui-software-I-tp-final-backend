package main

import (
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/app"
	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
