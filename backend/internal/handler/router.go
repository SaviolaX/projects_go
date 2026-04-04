package handler

import (
	"github.com/SaviolaX/blog/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(uh UserHandler, ph PostHandler, ch CategoryHandler, secret string) *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// User
	router.POST("/api/v1/auth/register", uh.Register)
	router.POST("/api/v1/auth/login", uh.Login)
	router.POST("/api/v1/auth/logout", uh.Logout)

	// Post
	router.GET("/api/v1/posts", ph.FindAll)
	router.GET("/api/v1/posts/:id", ph.FindByID)

	// Category
	router.GET("/api/v1/categories", ch.FindAll)

	// Secured
	postsGroup := router.Group("/api/v1/posts")

	postsGroup.Use(middleware.AuthMiddleware(secret))

	postsGroup.POST("/create", ph.Create)
	postsGroup.PUT("/:id", ph.Update)
	postsGroup.DELETE("/:id", ph.Delete)

	return router
}
