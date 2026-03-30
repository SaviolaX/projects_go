package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreatePage(ctx *gin.Context)
	UpdatePage(ctx *gin.Context)
	PostDetailPage(ctx *gin.Context)
	IndexPage(ctx *gin.Context)
}

type postHandler struct {
	service         service.PostService
	categoryService service.CategoryService
}

func (ph *postHandler) PostDetailPage(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	post, err := ph.service.FindByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	ctx.HTML(http.StatusOK, "detail_post.html", getTemplateData(ctx, gin.H{"post": post}))
}

func (ph *postHandler) CreatePage(ctx *gin.Context) {
	categories, err := ph.categoryService.FindAll()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}
	ctx.HTML(http.StatusOK, "create_post.html", gin.H{"categories": categories})
}

func (ph *postHandler) UpdatePage(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	post, err := ph.service.FindByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	log.Printf("Post category: %+v", post.Category)

	categories, err := ph.categoryService.FindAll()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.HTML(http.StatusOK, "update_post.html", getTemplateData(ctx, gin.H{"post": post, "categories": categories}))
}

func (ph *postHandler) IndexPage(ctx *gin.Context) {
	posts, err := ph.service.FindAll()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", getTemplateData(ctx, gin.H{
			"error": "Failed to load posts",
		}))
		return
	}

	ctx.HTML(http.StatusOK, "index.html", getTemplateData(ctx, gin.H{
		"posts": posts,
	}))
}

func (ph *postHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	err = ph.service.Delete(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (ph *postHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "invalid id"}))
		return
	}

	var updPostReq dto.UpdatePostRequest

	if ctx.PostForm("category_name") != "" {
		updPostReq.CategoryName = ctx.PostForm("category_name")
	} else {
		categoryIDStr := ctx.PostForm("category_id")
		if categoryIDStr != "" {
			categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
			if err != nil {
				ctx.Redirect(http.StatusSeeOther, "/posts/create?error=invalid category")
				return
			}
			updPostReq.CategoryID = uint(categoryID)
		}
	}

	updPostReq.Title = ctx.PostForm("title")
	updPostReq.Entry = ctx.PostForm("entry")

	err = ph.service.Update(uint(id), &updPostReq)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/posts/"+ctx.Param("id"))
}

func (ph *postHandler) Create(ctx *gin.Context) {
	var createReq dto.CreatePostRequest

	categoryIDStr := ctx.PostForm("category_id")
	if categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/posts/create?error=invalid category")
			return
		}
		createReq.CategoryID = uint(categoryID)
	}

	createReq.Title = ctx.PostForm("title")
	createReq.Entry = ctx.PostForm("entry")
	createReq.CategoryName = ctx.PostForm("category_name")

	err := createReq.Validate()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "don't have a userID"}))
		return
	}

	id, ok := userID.(uint)
	if !ok {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": "incorrect userID format"}))
		return
	}

	err = ph.service.Create(&createReq, id)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error.html", getTemplateData(ctx, gin.H{"error": err.Error()}))
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func NewPostHandler(service service.PostService, categoryService service.CategoryService) PostHandler {
	return &postHandler{
		service:         service,
		categoryService: categoryService,
	}
}
