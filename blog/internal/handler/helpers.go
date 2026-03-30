package handler

import "github.com/gin-gonic/gin"

func getTemplateData(ctx *gin.Context, data gin.H) gin.H {
	user, exists := ctx.Get("user")
	if exists {
		data["user"] = user
	}
	return data
}
