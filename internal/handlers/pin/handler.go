package pin

import (
	"Project1/internal/models"
	"Project1/internal/services/pin"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	PinService   pin.Service
}

func New(router *mux.Router, PinService pin.Service) {
	handler := &Handler{
		PinService: PinService,
	}
	p := router.PathPrefix("/pin").Subrouter()

	p.HandleFunc("/create", handler.Create).Methods("POST")
	p.HandleFunc("/get-by-userid", handler.GetByUserID).Methods("POST")
	p.HandleFunc("/get-by-id", handler.GetByID).Methods("POST")
	p.HandleFunc("/get-all", handler.GetAll).Methods("POST")
	p.HandleFunc("/delete", handler.Delete).Methods("POST")
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	b := models.Pin{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.PinService.Create(b)
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

	s, err := h.PinService.GetByUserID(id, 10, 10)
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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	Id, ok := mux.Vars(r)["id"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	id, err := strconv.Atoi(Id)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	s, err := h.PinService.GetByUserID(id, 10, 10)
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


func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	Id, ok := mux.Vars(r)["id"]
	if !ok{
		http.Error(w, " :( ", 400)
		return
	}

	id, err := strconv.Atoi(Id)
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}

	s, err := h.PinService.GetAll(id, 10)
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