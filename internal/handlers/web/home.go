package web

import (
	"Project1/internal/models"
	"Project1/internal/services/pin"
	"Project1/internal/services/users"
	"Project1/pkg/template"
	"context"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	router.HandleFunc("/create", handler.Create).Methods(http.MethodPost, http.MethodGet)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := h.UserService.GetByEmail(email)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if password == user.Password {
			http.SetCookie(w, &http.Cookie{
				Name:       "id",
				Value:      strconv.Itoa(user.ID),
				Expires:    time.Now().Add(24 * time.Hour),
			})
			http.Redirect(w,r,"/", http.StatusFound)
			return
		}

		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	files := []string{
		"./static/templates/sign-in.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("img")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		temp, err := ioutil.TempFile("static/img", "file_*.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer temp.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = temp.Write(fileBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.PinService.Create(models.Pin{
			Description:   "wadawdaw",
			AuthorID:      1,
			PinLink:       temp.Name(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w,r,"/", http.StatusFound)
	}

	files := []string{
		"./static/templates/create.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
