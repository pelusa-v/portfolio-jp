package services

import (
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/repositories"
)

type tagsService struct {
	repo repositories.TagsRepositoryPort // Service uses driven port
}

func NewTagsSevice(repo repositories.TagsRepositoryPort) *tagsService {
	return &tagsService{
		repo: repo,
	}
}

// Service implements driver port
func (srv *tagsService) UpdateTag(id int, tag entities.Tag) (entities.Tag, error) {
	return srv.repo.Update(id, tag)
}

func (srv *tagsService) DeleteTag(id int) error {
	return srv.repo.Delete(id)
}

func (srv *tagsService) CreateTag(tag entities.Tag) (entities.Tag, error) {
	return srv.repo.Create(tag)
}

func (srv *tagsService) GetTagsById(ids []uint) ([]entities.Tag, error) {
	return srv.GetTagsById(ids)
}
