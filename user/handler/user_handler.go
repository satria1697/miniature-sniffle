package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"six/middleware"
)

type UserHandler struct {
	r *mux.Router
}

func NewUserHandler(r *mux.Router) {
	handler := UserHandler{
		r: r,
	}
	r.Use(middleware.Authori)
	r.HandleFunc("/user", handler.GetUser)
}

func (h UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "konniciwa")
}
