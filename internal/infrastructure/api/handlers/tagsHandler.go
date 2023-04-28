package handlers

import (
	"fmt"
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

	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	handler.manager.CreatedSuccessfully(ctx, tag)
}

func (handler *tagsHandler) UpdateTag(ctx *gin.Context) {
	tagId, err := strconv.Atoi(ctx.Param("id"))
	// if err != nil {
	// 	handler.manager.BadRequest(ctx)
	// }

	var tagEntity entities.Tag
	ctx.BindJSON(&tagEntity) // load user value
	tag, err := handler.srv.UpdateTag(tagId, tagEntity)

	if err != nil {
		fmt.Print(err)
		handler.manager.BadRequest(ctx)
	} else {
		handler.manager.Success(ctx, tag)
	}
}

func (handler *tagsHandler) DeleteTag(ctx *gin.Context) {
	tagId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	err = handler.srv.DeleteTag(tagId)
	if err != nil {
		handler.manager.BadRequest(ctx)
	}

	handler.manager.Success(ctx, nil)
}
