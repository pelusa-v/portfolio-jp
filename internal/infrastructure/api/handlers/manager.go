package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Manager struct {
}

func (manager *Manager) Response(ctx *gin.Context, err error, object any, method string) {
	if err == nil {
		manager.ResponseSuccess(ctx, object, method)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
	} else if errors.Is(err, strconv.ErrSyntax) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	}
	return
}

func (manager *Manager) ResponseSuccess(ctx *gin.Context, object any, method string) {
	switch method {
	case "POST":
		ctx.AbortWithStatusJSON(http.StatusCreated, object)
		break
	case "PUT":
	case "DELETE":
	case "GET":
		fmt.Println(object)
		ctx.AbortWithStatusJSON(http.StatusOK, object)
		break
	}
}
