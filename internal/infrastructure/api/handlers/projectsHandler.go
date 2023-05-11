package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/app/requests"
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/services"
)

type projectsHandler struct {
	projectsSrv services.ProjectsServicePort
	tagsSrv     services.TagsServicePort
	manager     Manager
}

func NewProjectsHandler(projectsService services.ProjectsServicePort, tagsService services.TagsServicePort) *projectsHandler {
	return &projectsHandler{
		projectsSrv: projectsService,
		tagsSrv:     tagsService,
		manager:     Manager{},
	}
}

func (handler *projectsHandler) CreateProject(ctx *gin.Context) {
	var projectEntity entities.Project
	var projectRequest requests.CreateProjectRequest
	ctx.BindJSON(&projectRequest) // load project value

	tagsEntites, errTags := handler.tagsSrv.GetTagsById(projectRequest.Tags)
	if errTags != nil {
		ctx.AbortWithStatus(404)
	}

	projectEntity.Name = projectRequest.Name
	projectEntity.Summary = projectRequest.Summary
	projectEntity.Description = projectRequest.Description
	projectEntity.ImageURL = projectRequest.ImageURL
	projectEntity.Tags = tagsEntites

	project, err := handler.projectsSrv.CreateProject(projectEntity)

	handler.manager.Response(ctx, err, project, ctx.Request.Method)
}
