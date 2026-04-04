package handler

import (
	"net/http"

	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	FindAll(ctx *gin.Context)
}

type categoryHandler struct {
	service service.CategoryService
}

func (ch *categoryHandler) FindAll(ctx *gin.Context) {
	categories, err := ch.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"categories": categories})
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &categoryHandler{
		service: service,
	}
}
