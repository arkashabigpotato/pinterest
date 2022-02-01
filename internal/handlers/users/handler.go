package users

import (
	"Project1/internal/models"
	"Project1/internal/services/users"
	"encoding/json"
	"fmt"
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
	u.HandleFunc("/get-by-email/{email}", handler.GetByEmail).Methods("GET")
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	b := models.User{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(b)

	err = h.usersService.Create(b)
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
	email, ok := mux.Vars(r)["email"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	u, err := h.usersService.GetByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}