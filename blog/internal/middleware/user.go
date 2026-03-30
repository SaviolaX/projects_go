package middleware

import (
	"github.com/SaviolaX/blog/internal/auth"
	"github.com/SaviolaX/blog/internal/service"
	"github.com/gin-gonic/gin"
)

func UserMiddleware(userService service.UserService, secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, exists := ctx.Cookie("token")
		if exists != nil {
			ctx.Next()
			return
		}

		claims, err := auth.ValidateToken(token, secret)
		if err != nil {
			ctx.Next()
			return
		}

		user, err := userService.GetByID(claims.UserID)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
