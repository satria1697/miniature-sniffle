package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"six/middleware"
	"six/user/domain"
	util "six/utils"
)

type UserHandler struct {
	r           *mux.Router
	userService domain.UserService
}

func NewUserHandler(r *mux.Router, userService domain.UserService) {
	handler := UserHandler{
		r:           r,
		userService: userService,
	}
	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.Authori())
	s.HandleFunc("/user", handler.GetUser)
	s.HandleFunc("/user/{id}", handler.GetUserById)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	response := []domain.User{
		{
			ID:       1,
			Username: "bebek",
			Email:    "bebek@bebek.com",
		},
	}
	w.Write(util.Response(response, nil))
}

func (h UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	response, err := h.userService.GetUserById("2")
	if err != nil {
		w.Write(util.Response(nil, err))
		return
	}
	w.Write(util.Response(response, nil))
}
