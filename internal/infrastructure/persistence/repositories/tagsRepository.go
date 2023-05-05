package repositories

import (
	"errors"
	"fmt"

	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/config"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type tagsRepository struct {
	db *gorm.DB
}

func NewTagsRepository() *tagsRepository {
	db, _ := config.GetMysqlDB()
	return &tagsRepository{
		db: db,
	}
}

func (repo *tagsRepository) Update(id int, tag entities.Tag) (entities.Tag, error) {
	var tagToUpdate = models.Tag{}
	err := repo.db.First(&tagToUpdate, id).Error
	if err != nil {
		return tag, err
	}

	err = repo.db.Model(&tagToUpdate).Updates(tagToUpdate.MapEntityToModel(&tag)).Error // save execute update if model contains ID
	return tagToUpdate.MapModelToEntity(), err
}

func (repo *tagsRepository) Delete(id int) error {
	var tag models.Tag
	err := repo.db.First(&tag, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	repo.db.Delete(&tag)
	return nil
}

func (repo *tagsRepository) Create(tag entities.Tag) (entities.Tag, error) {
	var tagToCreate = models.Tag{}
	err := repo.db.Create(tagToCreate.MapEntityToModel(&tag)).Error
	fmt.Print(err)
	return tagToCreate.MapModelToEntity(), err
}
