package repository

import (
	"errors"
	"gorm.io/gorm"
	"six/auth/domain"
	userdomain "six/user/domain"
)

type authRepository struct {
	db *gorm.DB
}

func (a authRepository) LoginRepository(user userdomain.User) (string, error) {
	var password string
	res := a.db.Raw("SELECT password FROM users WHERE username = ?", user.Username)
	rows, _ := res.Rows()
	for rows.Next() {
		rows.Scan(&password)
	}
	if password == "" {
		return "", errors.New("email-not-found")
	}
	return password, nil
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return authRepository{
		db: db,
	}
}
