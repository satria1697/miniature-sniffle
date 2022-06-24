package handler

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"net/http"
	util "six/utils"
)

type AuthHandler struct {
	r *mux.Router
}

func NewAuthHandler(r *mux.Router) {
	handler := AuthHandler{
		r: r,
	}
	r.HandleFunc("/auth/login", handler.Login)
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
		tokenString, err := token.SignedString([]byte("inistringsecretkakak"))
		w.Header().Add("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write(util.Response(nil, err))
		}
		response := map[string]string{
			"token": tokenString,
		}
		w.Write(util.Response(response, nil))
	}
}
