package handler

import "github.com/gin-gonic/gin"

func SetupRouter(uh UserHandler, ph PostHandler) *gin.Engine {

	router := gin.Default()

	// User
	router.POST("/api/v1/auth/register", uh.Register)
	router.POST("/api/v1/auth/login", uh.Login)

	// Post
	router.GET("/api/v1/posts", ph.FindAll)
	router.GET("/api/v1/posts/:id", ph.FindByID)
	router.POST("/api/v1/posts", ph.Create)
	router.PUT("/api/v1/posts/:id", ph.Update)
	router.DELETE("/api/v1/posts/:id", ph.Delete)

	return router
}
