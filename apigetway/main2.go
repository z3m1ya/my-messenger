package main

/*
import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var serviceName = "API Gateway"

type Chat struct {
	ChatID   string    `json:"chatId"`
	Messages []Message `json:"messages"`
}

type Message struct {
	SenderID  string `json:"senderId"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type Profile struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Friends []string `json:"friends"`
}

type Response struct {
	Message string `json:"message"`
}

// GetChatsHandler обрабатывает запрос на получение списка чатов
func GetChatsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	friendID := r.URL.Query().Get("friendId")

	if userID == "" || friendID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "userId and friendId are required"})
		return
	}
	chats := []Chat{
		{
			ChatID: "1",
			Messages: []Message{
				{SenderID: "user1", Text: "Hello!", Timestamp: "2024-10-23T12:00:00Z"},
			},
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chats)
}

// AddMessageHandler обрабатывает запрос на добавление сообщения в чат
func AddMessageHandler(w http.ResponseWriter, r *http.Request) {
	chatID := mux.Vars(r)["chatId"]

	if chatID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "chatId is required"})
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request body"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "Message added successfully"})
}

// GetProfileHandler обрабатывает запрос на получение профиля по ID
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "id is required"})
		return
	}
	profile := Profile{
		ID:      id,
		Name:    "User " + id,
		Email:   "user" + id + "@example.com",
		Friends: []string{"friend1", "friend2"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

// UpdateProfileHandler обрабатывает запрос на изменение профиля
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "id is required"})
		return
	}
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request body"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Profile updated successfully"})
}

// RegisterHandler обрабатывает запрос на регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user Profile
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request body"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "User registered successfully"})
}

// LoginHandler обрабатывает запрос на авторизацию
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request body"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "User logged in successfully"})
}

// SendFriendRequestHandler обрабатывает запрос на отправку заявки в друзья
func SendFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	var friendRequest struct {
		SenderID   string `json:"senderId"`
		ReceiverID string `json:"receiverId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&friendRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request body"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "Friend request sent"})
}

// AcceptFriendRequestHandler обрабатывает запрос на принятие заявки в друзья
func AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	friendID := mux.Vars(r)["friendId"]

	if friendID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "friendId is required"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Friend request accepted"})
}

// UnfriendHandler обрабатывает запрос на удаление из друзей
func UnfriendHandler(w http.ResponseWriter, r *http.Request) {
	friendID := mux.Vars(r)["friendId"]

	if friendID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "friendId is required"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Friend removed"})
}

// GetFriendsHandler обрабатывает запрос на получение списка друзей
func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "userId is required"})
		return
	}

	friends := []string{"friend1", "friend2", "friend3"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/chats", GetChatsHandler).Methods("GET")
	router.HandleFunc("/chats/{chatId}/messages", AddMessageHandler).Methods("POST")
	router.HandleFunc("/profile/{id}", GetProfileHandler).Methods("GET")
	router.HandleFunc("/profile/{id}", UpdateProfileHandler).Methods("PUT")
	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/friends", SendFriendRequestHandler).Methods("POST")                    // Отправка запроса в друзья
	router.HandleFunc("/friends/{friendId}/accept", AcceptFriendRequestHandler).Methods("PUT") // Принятие заявки в друзья
	router.HandleFunc("/friends/{friendId}", UnfriendHandler).Methods("DELETE")                // Удаление из друзей
	router.HandleFunc("/friends/{userId}", GetFriendsHandler).Methods("GET")                   // Получение списка друзей

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("%s starting on port %s...", serviceName, "8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
*/
