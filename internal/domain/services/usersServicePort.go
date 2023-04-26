package services

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driver port
type UsersServicePort interface {
	GetUser(id int) (entities.User, error)
	ListUsers() ([]entities.User, error)
	UpdateUser(id int, user entities.User) (entities.User, error)
	DeleteUser(id int) error
	CreateUser(user entities.User) (entities.User, error)
	ValidatePasswordUser(user entities.User, psw string) (bool, error)
}
