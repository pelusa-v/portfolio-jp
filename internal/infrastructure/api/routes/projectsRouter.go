package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/app/services"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/api/handlers"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/repositories"
)

func RegisterProjectsRouter(router *gin.Engine) {
	// repos
	projectsRepository := repositories.NewProjectsRepository()
	tagsRepository := repositories.NewTagsRepository()
	// services
	projectsService := services.NewProjectsService(projectsRepository)
	tagsService := services.NewTagsSevice(tagsRepository)
	// handlers and routes
	projectsHandlers := handlers.NewProjectsHandler(projectsService, tagsService)
	router.POST("api/projects/add/", projectsHandlers.CreateProject)
	// router.GET("api/projects/:id", projectsHandlers.GetProject)
	// router.GET("api/projects/list/", projectsHandlers.ListProjects)
	// router.PUT("api/projects/update/:id", projectsHandlers.UpdateProject)
	// router.DELETE("api/projects/delete/:id", projectsHandlers.DeleteProject)
}
