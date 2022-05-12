package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ":)")
}

func signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ":( signUp")
}

func signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ":/ signIn")
}

func messages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ":0 semese")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-up", signUp)
	mux.HandleFunc("/sign-in", signIn)
	mux.HandleFunc("/messages", messages)

	http.ListenAndServe(":8080", mux)

}
