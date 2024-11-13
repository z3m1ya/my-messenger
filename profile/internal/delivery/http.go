package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"profile/internal/usecases"
)

type ProfileHandler struct {
	usecase usecases.ProfileUsecase
}

func NewProfileHandler(u usecases.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecase: u}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	profile, err := h.usecase.GetProfile(id)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updateRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.usecase.UpdateProfile(id, updateRequest.Name, updateRequest.Email)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
} //1

func (h *ProfileHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	friends, err := h.usecase.GetFriends(id)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friends)
}

func (h *ProfileHandler) SendFriendRequest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var request struct {
		TargetId string `json:"targetId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.usecase.SendFriendRequest(id, request.TargetId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
