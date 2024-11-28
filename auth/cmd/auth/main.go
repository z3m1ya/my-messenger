package main

import (
	"encoding/json"
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

func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// (здесь пример)
	if request.Token == "valid-token" {
		response := struct {
			Valid        bool   `json:"valid"`
			UserId       string `json:"userId"`
			ErrorMessage string `json:"errorMessage,omitempty"`
		}{
			Valid:  true,
			UserId: "12345",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		response := struct {
			Valid        bool   `json:"valid"`
			UserId       string `json:"userId,omitempty"`
			ErrorMessage string `json:"errorMessage"`
		}{
			Valid:        false,
			ErrorMessage: "Invalid token",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("User registered: %s", user.Username)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User registered successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if user.Username == "testuser" && user.Password == "password" { // Пример проверки
		response := struct {
			Message string `json:"message"`
			Token   string `json:"token"`
		}{
			Message: "Login successful",
			Token:   "valid-token",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health/liveness", livenessHandler).Methods("GET")
	r.HandleFunc("/health/readiness", readinessHandler).Methods("GET")

	r.HandleFunc("/api/auth/verify", verifyTokenHandler).Methods("POST")
	r.HandleFunc("/api/auth/register", registerHandler).Methods("POST")
	r.HandleFunc("/api/auth/login", loginHandler).Methods("POST")

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
