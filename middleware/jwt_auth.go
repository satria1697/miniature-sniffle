package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	util "six/utils"
	"strings"
)

func Authori() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(util.Response(nil, errors.New("unauthorized")))
				return
			}
			splitter := strings.Split(header, " ")
			if len(splitter) != 2 {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(util.Response(nil, errors.New("unauthorized")))
				return
			}
			token, err := jwt.Parse(splitter[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("error")
				}
				return []byte("inistringsecretkakak"), nil
			})
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(util.Response(nil, errors.New("unauthorized")))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if claims["email"].(string) != "" {
					handler.ServeHTTP(w, r)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(util.Response(nil, errors.New("unauthorized")))
				return
			} else {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(util.Response(nil, errors.New("unauthorized")))
				return
			}
		})
	}
}
