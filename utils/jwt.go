package util

import "github.com/golang-jwt/jwt/v4"

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, err := token.SignedString([]byte("inistringsecretkakak"))
	return tokenString, err
}
