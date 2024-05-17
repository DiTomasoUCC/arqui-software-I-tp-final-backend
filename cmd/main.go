package main

import (
	"log"

	"github.com/DiTomasoUCC/arqui-software-I-tp-final-backend/cmd/api"
)

func main() {
	server := api.NewAPIserver(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
