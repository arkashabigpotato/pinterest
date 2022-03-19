package message

import (
	"Project1/internal/models"
	"Project1/internal/services/message"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	MessageService   message.Service
}

func New(router *mux.Router, MessageService message.Service) {
	handler := &Handler{
		MessageService: MessageService,
	}
	m := router.PathPrefix("/message").Subrouter()

	m.HandleFunc("/create", handler.Create).Methods("POST")
	m.HandleFunc("/get", handler.Get).Methods("POST")
	m.HandleFunc("/delete", handler.Delete).Methods("POST")
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	b := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.MessageService.Create(b)
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	m := models.User{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil || m.ID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}

	mes, err := h.MessageService.Get(m.ID, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(mes)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request)  {
	a := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.MessageService.Delete(a.ID)
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