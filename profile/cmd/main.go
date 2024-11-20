package main

import (
	"log"
	"net/http"
	"profile/internal/delivery"
	"profile/internal/repo"
	"profile/internal/usecases"
	"time"

	"github.com/gorilla/mux"
)

var serviceName = "profile"

func main() {
	repo := repo.ProfileRepository{}
	usecase := usecases.NewProfileUsecase(repo)
	handler := delivery.NewProfileHandler(usecase)

	r := mux.NewRouter()
	r.HandleFunc("/profile/{id}", handler.GetProfile).Methods("GET")
	r.HandleFunc("/profile/{id}", handler.UpdateProfile).Methods("PUT")
	r.HandleFunc("/profile/{id}/friends", handler.GetFriends).Methods("GET")
	r.HandleFunc("/profile/{id}/friend-requests", handler.SendFriendRequest).Methods("POST")

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
