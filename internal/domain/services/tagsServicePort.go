package services

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driver port
type TagsServicePort interface {
	UpdateTag(id int, user entities.Tag) (entities.Tag, error)
	DeleteTag(id int) error
	CreateTag(tag entities.Tag) (entities.Tag, error)
	GetTagsById(ids []uint) ([]entities.Tag, error)
}
