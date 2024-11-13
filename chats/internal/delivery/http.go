package delivery

import (
	"encoding/json"
	"net/http"
	"time"

	"chats/api"
	"chats/internal/usecases"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HTTPHandler struct {
	useCase *usecases.ChatUseCase
}

func NewHTTPHandler(r *mux.Router, useCase *usecases.ChatUseCase) {
	handler := &HTTPHandler{useCase: useCase}

	r.HandleFunc("/chats", handler.listChatsHandler).Methods("GET")
	r.HandleFunc("/chats/{chatId}/messages", handler.addMessageHandler).Methods("POST")
}

func (h *HTTPHandler) authorizeToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", http.ErrNotSupported
	}

	if token != "valid-token" {
		return "", http.ErrNotSupported
	}

	return "userId", nil
}

func (h *HTTPHandler) listChatsHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := h.authorizeToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	friendId := r.URL.Query().Get("friendId")
	if friendId == "" {
		http.Error(w, "friendId is required", http.StatusBadRequest)
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

func (h *HTTPHandler) addMessageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := h.authorizeToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

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
