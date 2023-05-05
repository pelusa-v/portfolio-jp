package handlers

import (
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
	userId, _ := strconv.Atoi(ctx.Param("id"))
	userEntity, err := handler.srv.GetUser(userId)
	handler.manager.Response(ctx, err, userEntity, ctx.Request.Method)
}

func (handler *usersHandler) ListUsers(ctx *gin.Context) {
	usersEntities, err := handler.srv.ListUsers()
	handler.manager.Response(ctx, err, usersEntities, ctx.Request.Method)
}

func (handler *usersHandler) CreateUser(ctx *gin.Context) {
	var userEntity entities.User
	ctx.BindJSON(&userEntity) // load user value
	user, err := handler.srv.CreateUser(userEntity)

	handler.manager.Response(ctx, err, user, ctx.Request.Method)
}

func (handler *usersHandler) UpdateUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var userEntity entities.User
	ctx.BindJSON(&userEntity) // load user value
	user, err := handler.srv.UpdateUser(userId, userEntity)

	handler.manager.Response(ctx, err, user, ctx.Request.Method)
}

func (handler *usersHandler) DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	err := handler.srv.DeleteUser(userId)
	handler.manager.Response(ctx, err, nil, ctx.Request.Method)
}

func (handler *usersHandler) ValidateUserPassword(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var userValidation entities.UserValidation
	ctx.BindJSON(&userValidation)

	userEntity, err := handler.srv.GetUser(userId)

	if err != nil {
		handler.manager.Response(ctx, err, userEntity, ctx.Request.Method)
	} else {
		isValid, _ := handler.srv.ValidatePasswordUser(userEntity, userValidation.Password)
		handler.manager.Response(ctx, nil, gin.H{"validation": isValid}, ctx.Request.Method)
	}
}
