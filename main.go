package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"six/user/handler"
	"time"
)

func main() {
	r := mux.NewRouter()
	handler.NewUserHandler(r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
