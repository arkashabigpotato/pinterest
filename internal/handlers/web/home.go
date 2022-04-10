package web

import (
	"Project1/internal/services/pin"
	"Project1/pkg/template"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	PinService   pin.Service
}

func New(router *mux.Router, PinService pin.Service) {
	handler := &Handler{
		PinService: PinService,
	}

	router.HandleFunc("/", handler.Index)
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	files := []string{
		"./static/templates/index.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	pins, err := h.PinService.GetAll(100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"pins": pins,
	}

	err = template.ExecuteTemplate(ctx, w, files, data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
