package comment

import (
	"Project1/internal/models"
	"Project1/internal/services/comment"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	commentService   comment.Service
}

func New(router *mux.Router, commentService comment.Service) {
	handler := &Handler{
		commentService: commentService,
	}
	c := router.PathPrefix("/saved_pins").Subrouter()

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
	userId, ok := mux.Vars(r)["id"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	s, err := h.commentService.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userId, ok := mux.Vars(r)["id"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	s, err := h.commentService.GetByUserID(id, 10, 10)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) GetByPinID(w http.ResponseWriter, r *http.Request) {
	userId, ok := mux.Vars(r)["id"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	id, err := strconv.Atoi(userId)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	s, err := h.commentService.GetByPinID(id, 10, 10)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request)  {

}