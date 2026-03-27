package handler

import (
	"github.com/SaviolaX/blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(uh UserHandler, ph PostHandler, secret string) *gin.Engine {

	router := gin.Default()

	// User
	router.POST("/api/v1/auth/register", uh.Register)
	router.POST("/api/v1/auth/login", uh.Login)

	// Post
	router.GET("/api/v1/posts", ph.FindAll)
	router.GET("/api/v1/posts/:id", ph.FindByID)

	// Secured
	postsGroup := router.Group("/api/v1/posts")

	postsGroup.Use(middleware.AuthMiddleware(secret))

	postsGroup.POST("/", ph.Create)
	postsGroup.PUT("/:id", ph.Update)
	postsGroup.DELETE("/:id", ph.Delete)

	return router
}
