package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/app/services"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/api/handlers"
	"github.com/pelusa-v/portfolio-jp/internal/infrastructure/persistence/repositories"
)

func RegisterUsersRouter(router *gin.Engine) {
	usersRepository := repositories.NewUsersRepository()
	usersService := services.NewUsersSevice(usersRepository)
	usersHandlers := handlers.NewUsersHandler(usersService)
	router.GET("api/users/:id", usersHandlers.GetUser)
	router.GET("api/users/list/", usersHandlers.ListUsers)
	router.POST("api/users/add/", usersHandlers.CreateUser)
	router.PUT("api/users/update/:id", usersHandlers.UpdateUser)
	router.DELETE("api/users/delete/:id", usersHandlers.DeleteUser)
	router.POST("api/users/validate/:id", usersHandlers.ValidateUserPassword)
}
