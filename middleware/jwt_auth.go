package middleware

import (
	"fmt"
	"net/http"
)

func Authori(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		fmt.Printf("%v\n", header)
		if header == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}
