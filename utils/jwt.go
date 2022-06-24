package util

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Email: email,
	})
	tokenString, err := token.SignedString([]byte("inistringsecretkakak"))
	return tokenString, err
}
