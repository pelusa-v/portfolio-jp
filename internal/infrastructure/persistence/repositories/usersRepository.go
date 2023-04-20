package repositories

import (
	"errors"

	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/config"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository() *usersRepository {
	db, _ := config.GetMysqlDB()
	return &usersRepository{
		db: db,
	}
}

func (repo *usersRepository) Get(id int) (entities.User, error) {
	var user = models.User{}
	err := repo.db.First(&user, id).Error
	userEntity := user.MapModelToEntity()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userEntity, err
	}

	return userEntity, nil
}

// func (repo *usersRepository) List() ([]entities.User, error) {

// }

// func (repo *usersRepository) Update(id int) error {

// }

// func (repo *usersRepository) Delete(id int) error {

// }

// func (repo *usersRepository) Create(user entities.User) (entities.User, error) {

// }
