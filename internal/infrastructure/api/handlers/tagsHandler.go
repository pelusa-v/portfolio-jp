package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
	"github.com/pelusa-v/portfolio-jp/internal/domain/services"
)

type tagsHandler struct {
	srv     services.TagsServicePort
	manager Manager
}

func NewTagsHandler(service services.TagsServicePort) *tagsHandler {
	return &tagsHandler{
		srv:     service,
		manager: Manager{},
	}
}

func (handler *tagsHandler) CreateTag(ctx *gin.Context) {
	var tagEntity entities.Tag
	ctx.BindJSON(&tagEntity) // load user value
	tag, err := handler.srv.CreateTag(tagEntity)

	handler.manager.Response(ctx, err, tag, ctx.Request.Method)
}

func (handler *tagsHandler) UpdateTag(ctx *gin.Context) {
	tagId, _ := strconv.Atoi(ctx.Param("id"))
	var tagEntity entities.Tag
	ctx.BindJSON(&tagEntity) // load user value
	tag, err := handler.srv.UpdateTag(tagId, tagEntity)

	handler.manager.Response(ctx, err, tag, ctx.Request.Method)
}

func (handler *tagsHandler) DeleteTag(ctx *gin.Context) {
	tagId, _ := strconv.Atoi(ctx.Param("id"))
	err := handler.srv.DeleteTag(tagId)

	handler.manager.Response(ctx, err, nil, ctx.Request.Method)
}
