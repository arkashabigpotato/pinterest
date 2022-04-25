package users

import (
	"Project1/internal/models"
	"Project1/internal/services/users"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	usersService   users.Service
}

func New(router *mux.Router, usersService users.Service) {
	handler := &Handler{
		usersService: usersService,
	}

	u := router.PathPrefix("/users").Subrouter()

	u.HandleFunc("/create", handler.Create).Methods("POST")
	u.HandleFunc("/get-by-email", handler.GetByEmail).Methods("POST")
	u.HandleFunc("/get-by-id", handler.GetByID).Methods("POST")
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	b := models.User{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = h.usersService.Create(b)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode("ok")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil || u.Email == "" {
		http.Error(w, "Validation error", 400)
		return
	}

	user, err := h.usersService.GetByEmail(u.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil || u.ID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}


	us, err := h.usersService.GetByID(u.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(us)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}