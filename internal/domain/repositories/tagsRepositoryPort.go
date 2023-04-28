package repositories

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

// Driven port
type TagsRepositoryPort interface {
	Update(id int, tag entities.Tag) (entities.Tag, error)
	Delete(id int) error
	Create(tag entities.Tag) (entities.Tag, error)
}
