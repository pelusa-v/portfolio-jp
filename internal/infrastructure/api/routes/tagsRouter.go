package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/app/services"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/api/handlers"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/repositories"
)

func RegisterTagsRouter(router *gin.Engine) {
	tagRepository := repositories.NewTagsRepository()
	tagService := services.NewTagsSevice(tagRepository)
	tagHandlers := handlers.NewTagsHandler(tagService)
	router.POST("api/tags/add/", tagHandlers.CreateTag)
	router.PUT("api/tags/update/:id", tagHandlers.UpdateTag)
	router.DELETE("api/tags/delete/:id", tagHandlers.DeleteTag)
}
