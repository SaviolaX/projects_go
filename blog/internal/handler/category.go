package handler

import (
	"net/http"
	"strconv"

	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type categoryHandler struct {
	service service.CategoryService
}

func (ch *categoryHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	category, err := ch.service.FindByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.HTML(http.StatusOK, "category.html", getTemplateData(ctx, gin.H{"category": category}))
}

func (ch *categoryHandler) FindAll(ctx *gin.Context) {
	categories, err := ch.service.FindAll()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.HTML(http.StatusOK, "categories.html", getTemplateData(ctx, gin.H{"categories": categories}))
}

func (ch *categoryHandler) Create(ctx *gin.Context) {
	var createCategoryReq dto.CreateCategoryRequest

	createCategoryReq.Name = ctx.PostForm("name")

	err := createCategoryReq.Validate()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	err = ch.service.Create(&createCategoryReq)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (ch *categoryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	err = ch.service.Delete(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func NewCategoryHandler(service service.CategoryService) CategoryHandler {
	return &categoryHandler{
		service: service,
	}
}
