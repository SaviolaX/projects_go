package handler

import (
	"net/http"
	"strconv"

	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type postHandler struct {
	service service.PostService
}

func (ph *postHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = ph.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (ph *postHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	categoryID, err := strconv.ParseUint(ctx.PostForm("category_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
		return
	}

	var updPostReq dto.UpdatePostRequest

	updPostReq.Title = ctx.PostForm("title")
	updPostReq.Entry = ctx.PostForm("entry")
	updPostReq.CategoryID = uint(categoryID)

	err = ph.service.Update(uint(id), &updPostReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (ph *postHandler) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	post, err := ph.service.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

func (ph *postHandler) FindAll(ctx *gin.Context) {
	posts, err := ph.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (ph *postHandler) Create(ctx *gin.Context) {
	var createReq dto.CreatePostRequest

	createReq.Title = ctx.PostForm("title")
	createReq.Entry = ctx.PostForm("entry")

	err := createReq.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmpAuthorID := uint(1)

	err = ph.service.Create(&createReq, tmpAuthorID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func NewPostHandler(service service.PostService) PostHandler {
	return &postHandler{
		service: service,
	}
}
