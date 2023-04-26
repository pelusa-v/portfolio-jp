package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/services"
)

type usersHandler struct {
	srv     services.UsersServicePort
	manager Manager
}

func NewUsersHandler(service services.UsersServicePort) *usersHandler {
	return &usersHandler{
		srv:     service,
		manager: Manager{},
	}
}

func (handler *usersHandler) GetUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.manager.MakeGetResponse(ctx, err, nil)
	}

	userEntity, err := handler.srv.GetUser(userId)
	handler.manager.MakeGetResponse(ctx, err, userEntity)
}

func (handler *usersHandler) ListUsers(ctx *gin.Context) {
	usersEntities, err := handler.srv.ListUsers()
	handler.manager.MakeGetResponse(ctx, err, usersEntities)
}

func (handler *usersHandler) CreateUser(ctx *gin.Context) {
	var userEntity entities.User
	ctx.BindJSON(&userEntity) // load user value
	user, err := handler.srv.CreateUser(userEntity)

	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	handler.manager.CreatedSuccessfully(ctx, user)
}

func (handler *usersHandler) UpdateUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	handler.manager.BadRequest(ctx)
	// }

	var userEntity entities.User
	ctx.BindJSON(&userEntity) // load user value
	user, err := handler.srv.UpdateUser(userId, userEntity)

	if err != nil {
		fmt.Print(err)
		handler.manager.BadRequest(ctx)
	} else {
		handler.manager.Success(ctx, user)
	}
}

func (handler *usersHandler) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	err = handler.srv.DeleteUser(userId)
	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	handler.manager.Success(ctx, nil)
}
