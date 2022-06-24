package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"six/middleware"
	"six/user/domain"
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
	r.HandleFunc("/user/:id", handler.GetUser)
}

func (h UserHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	marshal, _ := json.Marshal(map[string][]domain.User{
		"data": {
			{
				ID:       1,
				Username: "bebek",
				Email:    "bebek@bebek.com",
			},
		},
	})
	writer.Write(marshal)
}

func (h UserHandler) GetUserById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	marshal, _ := json.Marshal(map[string]domain.User{
		"data": {
			ID:       1,
			Username: "bebek",
			Email:    "bebek@bebek.com",
		},
	})
	writer.Write(marshal)
}
