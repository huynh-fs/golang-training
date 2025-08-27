package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/huynh-fs/golang-training/user-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Lấy ID từ URL, ví dụ: /users/1
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}