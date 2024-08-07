package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getUserByID(w, r, id)
	case http.MethodPut:
		h.updateUser(w, r, id)
	case http.MethodDelete:
		h.deleteUser(w, r, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	user.ID = id
	if err := h.service.UpdateUser(&user); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.DeleteUser(id); err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
