package requests

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

type TagRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (tagReq *TagRequest) MapToEntity() entities.Tag {
	return entities.Tag{
		Name:        tagReq.Name,
		Description: tagReq.Description,
	}
}
