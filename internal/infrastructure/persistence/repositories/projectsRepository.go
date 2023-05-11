package repositories

import (
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/config"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/models"
	"gorm.io/gorm"
)

type projectsRepository struct {
	db *gorm.DB
}

func NewProjectsRepository() *projectsRepository {
	db, _ := config.GetMysqlDB()
	return &projectsRepository{
		db: db,
	}
}

func (repo *projectsRepository) Update(id int, project entities.Project) (entities.Project, error)
func (repo *projectsRepository) Delete(id int) error

func (repo *projectsRepository) Create(project entities.Project) (entities.Project, error) {
	var projectToCreate models.Project
	err := repo.db.Create(projectToCreate.MapEntityToModel(&project)).Error
	return projectToCreate.MapModelToEntity(), err
}

func (repo *projectsRepository) Get(id int) (entities.Project, error)
func (repo *projectsRepository) List() ([]entities.Project, error)
