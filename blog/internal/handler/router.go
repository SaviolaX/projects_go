package handler

import (
	"github.com/SaviolaX/blog/internal/middleware"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(uh UserHandler, ph PostHandler, ch CategoryHandler, us service.UserService, secret string) *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.Use(middleware.UserMiddleware(us, secret))

	// User
	router.GET("/login", uh.LoginPage)
	router.POST("/login", uh.LoginSubmit)
	router.GET("/register", uh.RegisterPage)
	router.POST("/register", uh.RegisterSubmit)
	router.POST("/logout", uh.Logout)

	// Post
	router.GET("/", ph.IndexPage)
	router.GET("/posts/:id", ph.PostDetailPage)
	router.GET("/posts/category/:categoryID", ph.PostsByCategory)

	router.GET("/categories", ch.FindAll)
	router.GET("/categories/:id", ch.FindByID)

	categoriesGroup := router.Group("/categories")

	categoriesGroup.Use(middleware.AuthMiddleware(secret))

	categoriesGroup.POST("/create", ch.Create)
	categoriesGroup.POST("/delete/:id", ch.Delete)

	// Secured
	postsGroup := router.Group("/posts")

	postsGroup.Use(middleware.AuthMiddleware(secret))

	postsGroup.POST("/create", ph.Create)
	postsGroup.GET("/create", ph.CreatePage)
	postsGroup.POST("/update/:id", ph.Update)
	postsGroup.GET("/update/:id", ph.UpdatePage)
	postsGroup.POST("/delete/:id", ph.Delete)

	return router
}
