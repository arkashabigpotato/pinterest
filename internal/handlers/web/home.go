package web

import (
	"Project1/internal/models"
	"Project1/internal/services/pin"
	"Project1/internal/services/users"
	"Project1/pkg/template"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	PinService   pin.Service
	UserService  users.Service
}

func New(router *mux.Router, PinService pin.Service, UserService  users.Service) {
	handler := &Handler{
		PinService: PinService,
		UserService:  UserService,
	}

	router.HandleFunc("/", handler.Index)
	router.HandleFunc("/profile", handler.Profile)
	router.HandleFunc("/sign-up", handler.SignUp).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/sign-in", handler.SignIn).Methods(http.MethodPost, http.MethodGet)
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

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	files := []string{
		"./static/templates/profile.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	user, err := h.UserService.GetByID(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pins, err := h.PinService.GetByUserID(1, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"pins": pins,
		"user": user,
	}

	err = template.ExecuteTemplate(ctx, w, files, data)
	if err != nil{
		fmt.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method == http.MethodPost {
		user := models.User{
			Email:      r.FormValue("email"),
			Password:   r.FormValue("password"),
			BirthDate:  r.FormValue("birth_date"),
			Username:   r.FormValue("username"),
			ProfileImg: "static/img/5.jpg",
			Status:     "qwertyuiop asdfghjkl zxcvbnm",
		}
		err := h.UserService.Create(user)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w,r,"/", http.StatusFound)
	}

	files := []string{
		"./static/templates/sign-up.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	data := map[string]interface{}{

	}

	err := template.ExecuteTemplate(ctx, w, files, data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	files := []string{
		"./static/templates/sign-in.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	data := map[string]interface{}{

	}

	err := template.ExecuteTemplate(ctx, w, files, data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
