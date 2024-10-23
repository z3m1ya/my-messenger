package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var serviceName = "auth"

// Liveness probe
func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Readiness probe
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ready"))
}

func main() {
	r := mux.NewRouter()

	// Определяем маршруты для проб
	r.HandleFunc("/health/liveness", livenessHandler).Methods("GET")
	r.HandleFunc("/health/readiness", readinessHandler).Methods("GET")

	// Устанавливаем параметры сервера
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
