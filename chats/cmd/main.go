package main

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
	"time"

	"chats/api"
	"github.com/gorilla/mux"
)

var serviceName = "chats"

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

func listChatsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	friendId := r.URL.Query().Get("friendId")

	if userId == "" || friendId == "" {
		http.Error(w, "userId and friendId are required", http.StatusBadRequest)
		return
	}

	sentTime := timestamppb.New(time.Now())
	response := &api.ListChatsResponse{
		Chats: []*api.Chat{
			{
				Id:           "chat1",
				Participants: []string{userId, friendId},
				Messages: []*api.Message{
					{
						Id:       "msg1",
						ChatId:   "chat1",
						SenderId: userId,
						Content:  "Hello!",
						SentTime: sentTime,
					},
				},
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	chatId := mux.Vars(r)["chatId"]

	var request api.AddMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if request.ChatId != chatId {
		http.Error(w, "Chat ID mismatch", http.StatusBadRequest)
		return
	}

	response := &api.AddMessageResponse{
		MessageId: "msg1",
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health/liveness", livenessHandler).Methods("GET")
	r.HandleFunc("/health/readiness", readinessHandler).Methods("GET")

	r.HandleFunc("/chats", listChatsHandler).Methods("GET")
	r.HandleFunc("/chats/{chatId}/messages", addMessageHandler).Methods("POST")

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
