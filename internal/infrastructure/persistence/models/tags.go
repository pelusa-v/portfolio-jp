package models

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

type Tag struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(200);not null"`
	Description string `gorm:"type:varchar(5000);not null"`
}

func (t *Tag) MapModelToEntity() entities.Tag {
	return entities.Tag{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
	}
}

func (t *Tag) MapEntityToModel(tag *entities.Tag) *Tag {
	t.ID = tag.ID
	t.Name = tag.Name
	t.Description = tag.Description
	return t
}
