package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"auth/internal/usecases"
)

type HTTPHandler struct {
	useCase *usecases.AuthUseCase
}

func NewHTTPHandler(r *mux.Router, useCase *usecases.AuthUseCase) {
	handler := &HTTPHandler{useCase: useCase}

	r.HandleFunc("/api/auth/verify", handler.verifyTokenHandler).Methods("POST")
	r.HandleFunc("/api/auth/register", handler.registerHandler).Methods("POST")
	r.HandleFunc("/api/auth/login", handler.loginHandler).Methods("POST")
}

func (h *HTTPHandler) verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Вызовите метод из usecase для верификации токена
	valid, userID, err := h.useCase.VerifyToken(request.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Valid  bool   `json:"valid"`
		UserId string `json:"userId,omitempty"`
	}{
		Valid:  valid,
		UserId: userID,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) registerHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.useCase.RegisterUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User registered successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.useCase.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Login successful",
		Token:   token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
