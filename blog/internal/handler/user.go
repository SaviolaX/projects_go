package handler

import (
	"net/http"

	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var loginReq dto.LoginRequest

	loginReq.Username = ctx.PostForm("username")
	loginReq.Password = ctx.PostForm("password")

	err := loginReq.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.service.Login(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var registerReq dto.RegisterRequest

	registerReq.Username = ctx.PostForm("username")
	registerReq.Email = ctx.PostForm("email")
	registerReq.Password = ctx.PostForm("password")

	err := registerReq.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uh.service.Register(&registerReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{
		service: service,
	}
}
