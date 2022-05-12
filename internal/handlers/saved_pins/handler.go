package saved_pins

import (
	"Project1/internal/models"
	"Project1/internal/services/saved_pins"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	savedPinService saved_pins.Service
}

func New(router *mux.Router, savedPinService saved_pins.Service) {
	handler := &Handler{
		savedPinService: savedPinService,
	}
	s := router.PathPrefix("/saved_pins").Subrouter()

	s.HandleFunc("/append", handler.Append).Methods("POST")
	s.HandleFunc("/get-by-userid", handler.GetByUserID).Methods("POST")
	s.HandleFunc("/delete", handler.Delete).Methods("POST")
}

func (h *Handler) Append(w http.ResponseWriter, r *http.Request) {
	b := models.SavedPin{}

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.savedPinService.Append(b)
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
	sp := models.SavedPin{}

	err := json.NewDecoder(r.Body).Decode(&sp)
	if err != nil || sp.UserID == 0 {
		http.Error(w, "Validation error", http.StatusBadRequest)
		return
	}

	s, err := h.savedPinService.GetByUserID(sp.UserID, 10, 0)
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

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	a := models.SavedPin{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = h.savedPinService.Delete(a.PinID)
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
