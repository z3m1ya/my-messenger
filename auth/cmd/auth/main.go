package main

import (
	"log"
	"net/http"
	"time"

	"auth/internal/delivery"
	"auth/internal/usecases"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	authUseCase := usecases.NewAuthUseCase()

	delivery.NewHTTPHandler(r, authUseCase)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
