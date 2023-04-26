package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Manager struct {
}

func (manager *Manager) MakeGetResponse(ctx *gin.Context, err error, object any) {
	if err == nil {
		ctx.JSON(http.StatusOK, object)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
	} else if errors.Is(err, strconv.ErrSyntax) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	}
	return
}

func (manager *Manager) BadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	return
}

func (manager *Manager) CreatedSuccessfully(ctx *gin.Context, object any) {
	ctx.JSON(http.StatusCreated, object)
	return
}

func (manager *Manager) Success(ctx *gin.Context, object any) {
	ctx.JSON(http.StatusOK, object)
	return
}
