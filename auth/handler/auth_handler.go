package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"six/auth/domain"
	userdomain "six/user/domain"
	util "six/utils"
)

type AuthHandler struct {
	r           *mux.Router
	authService domain.AuthService
}

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(r *mux.Router, authService domain.AuthService) {
	handler := AuthHandler{
		r:           r,
		authService: authService,
	}
	r.HandleFunc("/auth/login", handler.Login)
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		all, err := ioutil.ReadAll(r.Body)
		var request AuthLoginRequest
		json.Unmarshal(all, &request)
		res, err := h.authService.LoginService(userdomain.User{
			Username: request.Username,
			Password: request.Password,
		})
		w.Header().Add("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write(util.Response(nil, err))
			return
		}
		response := map[string]string{
			"token": res,
		}
		w.Write(util.Response(response, nil))
	}
}
