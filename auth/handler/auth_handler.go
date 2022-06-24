package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"net/http"
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
	fmt.Printf("%s\n", r.Method)
	if r.Method == "POST" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
		fmt.Printf("%v\n", token.Raw)
		tokenString, err := token.SignedString([]byte("inistringsecretkakak"))
		w.Header().Add("Content-Type", "application/json")

		if err != nil {
			fmt.Fprintf(w, "%v", err)
		}
		fmt.Fprintf(w, "%v", tokenString)
	}
}
