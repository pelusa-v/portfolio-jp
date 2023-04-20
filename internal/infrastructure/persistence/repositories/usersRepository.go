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

func (repo *usersRepository) List() ([]entities.User, error) {
	var users models.ListUser
	repo.db.Find(&users)
	return users.MapModelToEntity(), nil
}

func (repo *usersRepository) Update(id int, user entities.User) error {
	err := repo.db.Save(user).Error // save execute update if model contains ID
	return err
}

func (repo *usersRepository) Delete(id int) error {
	var user models.User
	err := repo.db.Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	repo.db.Delete(user)
	return nil
}

func (repo *usersRepository) Create(user entities.User) (entities.User, error) {
	err := repo.db.Create(models.MapEntityToModel(&user)).Error
	return user, err
}
