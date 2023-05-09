package models

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(200);uniqueIndex"`
	Summary     string `gorm:"type:varchar(500)"`
	Description string `gorm:"type:varchar(5000)"`
	ImageURL    string `gorm:"type:varchar(500)"`
	Tags        []Tag  `gorm:"many2many:project_tags;"` // project_tags is the join table
	// References  []Reference `json:"URL"`
}

func (project *Project) MapModelToEntity() *entities.Project {
	tagsEntities := TagList(project.Tags)
	return &entities.Project{
		ID:          project.ID,
		Name:        project.Name,
		Summary:     project.Summary,
		Description: project.Description,
		ImageURL:    project.ImageURL,
		Tags:        tagsEntities.MapModelToEntity(),
	}
}
