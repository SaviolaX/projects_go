package handler

import (
	"net/http"
	"strings"

	"github.com/SaviolaX/blog/internal/auth"
	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userHandler struct {
	expiredHours int
	secret       string
	service      service.UserService
}

func (uh *userHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := auth.ValidateToken(tokenString, uh.secret)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "logged out"})
		return
	}

	auth.Blocklist.Add(tokenString, claims.ExpiresAt.Time)

	ctx.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var loginReq dto.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect reques body"})
	}

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

	jwtToken, err := auth.GenerateToken(user.ID, uh.secret, uh.expiredHours)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "jwt token is not generated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
		"token":     jwtToken,
	})
}

func (uh *userHandler) Register(ctx *gin.Context) {
	var registerReq dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect reques body"})
	}

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

func NewUserHandler(expiredHours int, secret string, service service.UserService) UserHandler {
	return &userHandler{
		expiredHours: expiredHours,
		secret:       secret,
		service:      service,
	}
}
