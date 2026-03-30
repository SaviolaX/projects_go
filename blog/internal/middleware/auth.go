package middleware

import (
	"net/http"

	"github.com/SaviolaX/blog/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader, err := ctx.Cookie("token")
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(authHeader, secret)
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}
