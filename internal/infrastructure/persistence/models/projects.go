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

func (project *Project) MapModelToEntity() entities.Project {
	tagsEntities := TagList(project.Tags)
	return entities.Project{
		ID:          project.ID,
		Name:        project.Name,
		Summary:     project.Summary,
		Description: project.Description,
		ImageURL:    project.ImageURL,
		Tags:        tagsEntities.MapModelToEntity(),
	}
}

func (proj *Project) MapEntityToModel(project *entities.Project) Project {
	var tags []Tag

	for _, tagEntity := range project.Tags {
		var tagModel Tag
		tags = append(tags, *tagModel.MapEntityToModel(&tagEntity))
	}

	proj.ID = project.ID
	proj.Name = project.Name
	proj.Description = project.Description
	proj.ImageURL = project.ImageURL
	proj.Tags = tags
	return *proj
}
