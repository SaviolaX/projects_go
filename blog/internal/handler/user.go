package handler

import (
	"net/http"

	"github.com/SaviolaX/blog/internal/auth"
	"github.com/SaviolaX/blog/internal/dto"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	LoginPage(ctx *gin.Context)
	LoginSubmit(ctx *gin.Context)
	RegisterPage(ctx *gin.Context)
	RegisterSubmit(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userHandler struct {
	expiredHours int
	secret       string
	service      service.UserService
}

func (uh *userHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}

func (uh *userHandler) RegisterSubmit(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if exists {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	var registerReq dto.RegisterRequest

	registerReq.Username = ctx.PostForm("username")
	registerReq.Email = ctx.PostForm("email")
	registerReq.Password = ctx.PostForm("password")

	err := registerReq.Validate()
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/register?error="+err.Error())
		return
	}

	err = uh.service.Register(&registerReq)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/register?error="+err.Error())
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/login")
}

func (uh *userHandler) RegisterPage(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if exists {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.HTML(http.StatusOK, "register_user.html", getTemplateData(ctx, gin.H{"error": ctx.Query("error")}))
}

func (uh *userHandler) LoginSubmit(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if exists {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	var loginReq dto.LoginRequest

	loginReq.Username = ctx.PostForm("username")
	loginReq.Password = ctx.PostForm("password")

	err := loginReq.Validate()

	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	user, err := uh.service.Login(&loginReq)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	jwtToken, err := auth.GenerateToken(user.ID, uh.secret, uh.expiredHours)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login?error="+err.Error())
		return
	}

	ctx.SetCookie("token", jwtToken, uh.expiredHours*3600, "/", "", false, true)
	ctx.Redirect(http.StatusSeeOther, "/")
}

func (uh *userHandler) LoginPage(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if exists {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.HTML(http.StatusOK, "login_user.html", getTemplateData(ctx, gin.H{"error": ctx.Query("error")}))
}

func NewUserHandler(expiredHours int, secret string, service service.UserService) UserHandler {
	return &userHandler{
		expiredHours: expiredHours,
		secret:       secret,
		service:      service,
	}
}
