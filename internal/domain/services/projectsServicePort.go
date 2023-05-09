package services

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driver port
type ProjectsServicePort interface {
	UpdateProject(id int, project entities.Project) (entities.Project, error)
	DeleteProject(id int) error
	CreateProject(project entities.Project) (entities.Project, error)
	GetProject(id int) (entities.Project, error)
	ListProjects() ([]entities.Project, error)
}
