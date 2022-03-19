package comment

import (
	"Project1/internal/models"
	"Project1/internal/services/comment"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	commentService   comment.Service
}

func New(router *mux.Router, commentService comment.Service) {
	handler := &Handler{
		commentService: commentService,
	}
	c := router.PathPrefix("/comment").Subrouter()

	c.HandleFunc("/create", handler.Create).Methods("POST")
	c.HandleFunc("/get-by-id", handler.GetByID).Methods("POST")
	c.HandleFunc("/get-by-userid", handler.GetByUserID).Methods("POST")
	c.HandleFunc("/get-by-pinid", handler.GetByPinID).Methods("POST")
	c.HandleFunc("/delete", handler.Delete).Methods("POST")
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	b := models.Comment{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.commentService.Create(b)
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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	c := models.Comment{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil || c.ID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}


	com, err := h.commentService.GetByID(c.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(com)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	c := models.Comment{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil || c.AuthorID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}


	com, err := h.commentService.GetByUserID(c.AuthorID, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(com)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByPinID(w http.ResponseWriter, r *http.Request) {
	c := models.Comment{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil || c.PinID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}


	com, err := h.commentService.GetByPinID(c.PinID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(com)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request)  {
	a := models.Comment{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.commentService.Delete(a.ID)
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