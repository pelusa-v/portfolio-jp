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
	return userEntity, err
}

func (repo *usersRepository) List() ([]entities.User, error) {
	var users models.ListUser
	repo.db.Find(&users)
	return users.MapModelToEntity(), nil
}

func (repo *usersRepository) Update(id int, user entities.User) (entities.User, error) {
	var userToUpdate = models.User{}
	err := repo.db.First(&userToUpdate, id).Error
	if err != nil {
		return user, err
	}

	err = repo.db.Model(&userToUpdate).Updates(userToUpdate.MapEntityToModel(&user)).Error // save execute update if model contains ID
	return user, err
}

func (repo *usersRepository) Delete(id int) error {
	var user models.User
	err := repo.db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	repo.db.Delete(&user)
	return nil
}

func (repo *usersRepository) Create(user entities.User) (entities.User, error) {
	var userToCreate = models.User{}
	err := repo.db.Create(userToCreate.MapEntityToModel(&user)).Error
	return user, err
}
