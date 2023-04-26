package repositories

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driven port
type UsersRepositoryPort interface {
	Get(id int) (entities.User, error)
	List() ([]entities.User, error)
	Update(id int, user entities.User) (entities.User, error)
	Delete(id int) error
	Create(user entities.User) (entities.User, error)
}
