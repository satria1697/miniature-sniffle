package middleware

import (
	"fmt"
	"net/http"
)

func Authori(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		fmt.Printf("%v\n", header)
		handler.ServeHTTP(w, r)
	})
}
