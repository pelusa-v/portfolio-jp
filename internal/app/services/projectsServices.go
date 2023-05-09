package services

import (
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/repositories"
)

type projectsService struct {
	repository repositories.ProjectsRepositoryPort
}

func NewProjectsService(repo repositories.ProjectsRepositoryPort) *projectsService {
	return &projectsService{
		repository: repo,
	}
}

func (srv *projectsService) UpdateProject(id int, project entities.Project) (entities.Project, error) {
	return srv.repository.Update(id, project)
}

func (srv *projectsService) DeleteProject(id int) error {
	return srv.repository.Delete(id)
}

func (srv *projectsService) CreateProject(project entities.Project) (entities.Project, error) {
	return srv.repository.Create(project)
}

func (srv *projectsService) GetProject(id int) (entities.Project, error) {
	return srv.repository.Get(id)
}

func (srv *projectsService) ListProjects() ([]entities.Project, error) {
	return srv.repository.List()
}
