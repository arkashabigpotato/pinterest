package web

import (
	"Project1/internal/models"
	"Project1/internal/services/comment"
	"Project1/internal/services/message"
	"Project1/internal/services/pin"
	"Project1/internal/services/saved_pins"
	"Project1/internal/services/users"
	"Project1/pkg/ctx_data"
	"Project1/pkg/template"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	PinService       pin.Service
	UserService      users.Service
	SavedPinsService saved_pins.Service
	CommentService   comment.Service
	MessageService   message.Service
}

func New(router *mux.Router, PinService pin.Service, UserService users.Service, SavedPinsService saved_pins.Service, CommentService comment.Service, MessageService message.Service) {
	handler := &Handler{
		PinService:       PinService,
		UserService:      UserService,
		SavedPinsService: SavedPinsService,
		CommentService:   CommentService,
		MessageService:   MessageService,
	}

	router.HandleFunc("/", handler.Index)
	router.HandleFunc("/profile", handler.Profile)
	router.HandleFunc("/create", handler.Create).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/sign-up", handler.SignUp).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/sign-in", handler.SignIn).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/logout", handler.Logout)
	router.HandleFunc("/saved-pins", handler.SavedPins)
	router.HandleFunc("/pin/{id}", handler.PinPage).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/save/{id}", handler.Save).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/like/{id}", handler.Like).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/dislike/{id}", handler.Dislike).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/settings", handler.Settings).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/messages/{id}", handler.Message).Methods(http.MethodPost, http.MethodGet)
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	isLoggedIn := true
	if userData.UserID == 0 {
		isLoggedIn = false
	}

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
		"pins":       pins,
		"isLoggedIn": isLoggedIn,
	}

	err = template.ExecuteTemplate(ctx, w, files, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/profile.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	user, err := h.UserService.GetByID(userData.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pins, err := h.PinService.GetByUserID(userData.UserID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"pins":       pins,
		"user":       user,
		"isLoggedIn": true,
	}

	err = template.ExecuteTemplate(r.Context(), w, files, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		user := models.User{
			Email:      r.FormValue("email"),
			Password:   r.FormValue("password"),
			BirthDate:  r.FormValue("birth_date"),
			Username:   r.FormValue("username"),
			ProfileImg: "static/img/5.jpg",
			Status:     "qwertyuiop asdfghjkl zxcvbnm",
		}
		id, err := h.UserService.Create(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "id",
			Value:   strconv.Itoa(id),
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/sign-up.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{"isLoggedIn": false})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID != 0 {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := h.UserService.GetByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if password == user.Password {
			http.SetCookie(w, &http.Cookie{
				Name:    "id",
				Value:   strconv.Itoa(user.ID),
				Expires: time.Now().Add(24 * time.Hour),
			})
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		http.Error(w, "bad password", http.StatusBadRequest)
		return
	}

	files := []string{
		"./static/templates/sign-in.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{"isLoggedIn": false})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		description := r.FormValue("description")
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
			Description: description,
			AuthorID:    userData.UserID,
			PinLink:     temp.Name(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/create.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{"isLoggedIn": true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "id",
		Expires: time.Now().Add(-1 * time.Hour),
	})
	http.Redirect(w, r.WithContext(context.Background()), "/", http.StatusFound)
}

func (h *Handler) SavedPins(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	savedPins, err := h.SavedPinsService.GetByUserID(userData.UserID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var pins []*models.Pin
	for _, savedPin := range savedPins {
		p, err := h.PinService.GetByID(savedPin.PinID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pins = append(pins, p)
	}
	data := map[string]interface{}{
		"pins":       pins,
		"isLoggedIn": true,
	}

	files := []string{
		"./static/templates/saved-pins.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err = template.ExecuteTemplate(ctx, w, files, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) PinPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)
	vars := mux.Vars(r)

	pinID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("message")
		err := h.CommentService.Create(models.Comment{
			PinID:    pinID,
			Text:     text,
			AuthorID: userData.UserID,
			DateTime: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/pin/"+strconv.Itoa(pinID), http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/pin-page.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	p, err := h.PinService.GetByID(pinID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	com, err := h.CommentService.GetByPinID(pinID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"pin":        p,
		"isLoggedIn": true,
		"comments":   com,
	}

	err = template.ExecuteTemplate(r.Context(), w, files, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pinID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	err = h.SavedPinsService.Append(models.SavedPin{
		PinID:  pinID,
		UserID: userData.UserID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pin/"+vars["id"], http.StatusFound)
	return
}

func (h *Handler) Like(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pinID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	err = h.PinService.Like(pinID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pin/"+vars["id"], http.StatusFound)
	return
}

func (h *Handler) Dislike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pinID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	err = h.PinService.Dislike(pinID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pin/"+vars["id"], http.StatusFound)
	return
}

func (h *Handler) Settings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)

	if userData.UserID == 0 {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost {
		status := r.FormValue("status")
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

		err = h.UserService.Update(models.User{
			ID:         userData.UserID,
			ProfileImg: temp.Name(),
			Status:     status,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/settings.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	err := template.ExecuteTemplate(ctx, w, files, map[string]interface{}{"isLoggedIn": true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Message(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userData := ctx_data.FromContext(ctx)
	vars := mux.Vars(r)

	chatID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("message")
		err := h.MessageService.Create(models.Message{
			FromID:   userData.UserID,
			ToID:     chatID,
			Text:     text,
			DateTime: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/messages/"+strconv.Itoa(chatID), http.StatusFound)
		return
	}

	files := []string{
		"./static/templates/messages.page.tmpl",
		"./static/templates/base.layout.tmpl",
	}

	m, err := h.MessageService.Get(userData.UserID, chatID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"isLoggedIn": true,
		"messages":   m,
	}

	err = template.ExecuteTemplate(r.Context(), w, files, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
