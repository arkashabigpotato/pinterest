package middleware

import (
	"Project1/pkg/ctx_data"
	"log"
	"net/http"
	"strconv"
)

func LoggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI, r.Method, r.Body, r.UserAgent())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func ContextDataMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookie, err := r.Cookie("id")
		if err != nil && err == http.ErrNoCookie {
			ctx = ctx_data.ToContext(ctx, ctx_data.UserData{UserID: 0})
			next.ServeHTTP(w, r.Clone(ctx).WithContext(ctx))
			return
		}
		if err != nil && err != http.ErrNoCookie{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx = ctx_data.ToContext(ctx, ctx_data.UserData{UserID: id})
		next.ServeHTTP(w, r.Clone(ctx).WithContext(ctx))
	})
}