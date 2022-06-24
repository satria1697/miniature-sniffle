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

func (a authRepository) LoginRepository(user userdomain.User) (string, string, error) {
	var data userdomain.User
	res := a.db.Raw("SELECT email, password FROM users WHERE username = ?", user.Username)
	rows, _ := res.Rows()
	for rows.Next() {
		rows.Scan(&data.Email, &data.Password)
	}
	if data.Password == "" || data.Email == "" {
		return "", "", errors.New("not-found-cred")
	}
	return data.Email, data.Password, nil
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return authRepository{
		db: db,
	}
}
