package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"six/auth/domain"
	userdomain "six/user/domain"
	util "six/utils"
)

type authService struct {
	authRepository domain.AuthRepository
}

func checkPassword(password string, dbPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (a authService) LoginService(user userdomain.User) (string, error) {
	email, password, err := a.authRepository.LoginRepository(user)
	if err != nil {
		return "", err
	}

	err = checkPassword(user.Password, password)
	if err != nil {
		return "", errors.New("email-not-found")
	}

	return util.GenerateToken(email)
}

func NewAuthService(authRepository domain.AuthRepository) domain.AuthService {
	return authService{
		authRepository: authRepository,
	}
}
