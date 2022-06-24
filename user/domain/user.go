package domain

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	GetUserById(id string) (User, error)
}
