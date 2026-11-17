package main

import (
	"log"
	"os"

	"github.com/szymczykkrzysztof/rachunki/backend/internal/httpapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := httpapi.NewRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
