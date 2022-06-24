package domain

import "six/user/domain"

type AuthService interface {
	LoginService(user domain.User) (string, error)
}

type AuthRepository interface {
	LoginRepository(user domain.User) (string, error)
}
