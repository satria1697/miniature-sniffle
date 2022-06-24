package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"six/auth/domain"
	userdomain "six/user/domain"
	util "six/utils"
)

type authService struct {
	authRepository domain.AuthRepository
}

func checkPassword(password string, dbPassword string) error {
	fmt.Printf("%v\n", dbPassword)
	fmt.Printf("%v\n", password)
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (a authService) LoginService(user userdomain.User) (string, error) {
	res, err := a.authRepository.LoginRepository(user)
	if err != nil {
		return "", err
	}

	err = checkPassword(user.Password, res)
	if err != nil {
		return "", errors.New("email-not-found")
	}

	return util.GenerateToken(res)
}

func NewAuthService(authRepository domain.AuthRepository) domain.AuthService {
	return authService{
		authRepository: authRepository,
	}
}
