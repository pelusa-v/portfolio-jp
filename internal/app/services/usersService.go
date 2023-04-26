package services

import (
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/repositories"
)

type usersService struct {
	repo repositories.UsersRepositoryPort // Service uses driven port
}

func NewUsersSevice(repo repositories.UsersRepositoryPort) *usersService {
	return &usersService{
		repo: repo,
	}
}

// Service implements driver port
func (srv *usersService) GetUser(id int) (entities.User, error) {
	return srv.repo.Get(id)
}

func (srv *usersService) ListUsers() ([]entities.User, error) {
	return srv.repo.List()
}

func (srv *usersService) UpdateUser(id int, user entities.User) (entities.User, error) {
	return srv.repo.Update(id, user)
}

func (srv *usersService) DeleteUser(id int) error {
	return srv.repo.Delete(id)
}

func (srv *usersService) CreateUser(user entities.User) (entities.User, error) {
	return srv.repo.Create(user)
}

func (srv *usersService) ValidatePasswordUser(user entities.User, psw string) (bool, error) {
	return false, nil // TODO implement this function
}
