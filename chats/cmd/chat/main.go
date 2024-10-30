package main

import (
	"log"
	"net/http"
	"time"

	"chats/internal/delivery"
	"chats/internal/repo"
	"chats/internal/usecases"
	"github.com/gorilla/mux"
)

var serviceName = "chats"

func main() {
	r := mux.NewRouter()

	// Инициализация репозиториев и usecases
	userRepo := repo.NewInMemoryUserRepository()
	chatUseCase := usecases.NewChatUseCase(userRepo)

	// Инициализация HTTP-обработчиков
	delivery.NewHTTPHandler(r, chatUseCase)

	// Здоровье проверки
	r.HandleFunc("/health/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	r.HandleFunc("/health/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready"))
	}).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("%s starting on port %s...", serviceName, "8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
