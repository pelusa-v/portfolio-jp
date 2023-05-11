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

type TagList []Tag

func (tagList *TagList) MapModelToEntity() []entities.Tag {
	var tagsEntities []entities.Tag
	for _, tag := range *tagList {
		tagsEntities = append(tagsEntities, tag.MapModelToEntity())
	}

	return tagsEntities
}

// func (tagList *TagList) MapEntityToModel(tags *[]entities.Tag) []Tag {
// 	for _, tagEntity := range *tags {
// 		tagList = append(tagList, tagEntity.)
// 	}

// 	t.ID = tag.ID
// 	t.Name = tag.Name
// 	t.Description = tag.Description
// 	return []Tag(*tagList)
// }
