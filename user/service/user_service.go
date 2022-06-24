package service

import (
	"errors"
	"fmt"
	"six/user/domain"
	"strconv"
)

type userService struct {
}

func (u userService) GetUserById(id string) (domain.User, error) {
	idInt, _ := strconv.ParseInt(id, 10, 64)
	fmt.Printf("%v\n", idInt)
	if idInt > 1 {
		return domain.User{}, errors.New("NO ID FOUND")
	}
	return domain.User{
		ID:       1,
		Username: "bebek",
		Email:    "bebek@bebek.com",
	}, nil
}

func NewUserService() domain.UserService {
	return userService{}
}
