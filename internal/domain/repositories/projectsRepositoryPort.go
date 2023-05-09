package repositories

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driven port
type ProjectsRepositoryPort interface {
	Update(id int, project entities.Project) (entities.Project, error)
	Delete(id int) error
	Create(project entities.Project) (entities.Project, error)
	Get(id int) (entities.Project, error)
	List() ([]entities.Project, error)
}
