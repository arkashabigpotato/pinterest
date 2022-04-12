package main

import (
	comment3 "Project1/internal/handlers/comment"
	message3 "Project1/internal/handlers/message"
	pin3 "Project1/internal/handlers/pin"
	saved_pins3 "Project1/internal/handlers/saved_pins"
	users3 "Project1/internal/handlers/users"
	"Project1/internal/handlers/web"
	"Project1/internal/repository/comment"
	"Project1/internal/repository/message"
	"Project1/internal/repository/pin"
	"Project1/internal/repository/saved_pins"
	"Project1/internal/repository/users"
	comment2 "Project1/internal/services/comment"
	message2 "Project1/internal/services/message"
	pin2 "Project1/internal/services/pin"
	saved_pins2 "Project1/internal/services/saved_pins"
	users2 "Project1/internal/services/users"
	db2 "Project1/pkg/db"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := db2.PostgresConnection()
	router := mux.NewRouter()
	userRepo := users.NewUserRepository(db)
	userService := users2.NewService(userRepo)
	users3.New(router, userService)

	savedPinRepo := saved_pins.NewSavedPinRepository(db)
	savedPinService := saved_pins2.NewService(savedPinRepo)
	saved_pins3.New(router, savedPinService)

	pinRepo := pin.NewPinRepository(db)
	pinService := pin2.NewService(pinRepo)
	pin3.New(router, pinService)

	messageRepo := message.NewMessageRepository(db)
	messageService := message2.NewService(messageRepo)
	message3.New(router, messageService)

	commentRepo := comment.NewCommentRepository(db)
	commentService := comment2.NewService(commentRepo)
	comment3.New(router, commentService)

	web.New(router, pinService, userService)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe(":145", router)
	if err != nil {
		fmt.Println(err)
	}
}
