package middleware

import (
	"net/http"
	"strings"

	"github.com/SaviolaX/blog/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ValidateToken(tokenString, secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		ctx.Set("userID", claims.UserID)

		ctx.Next()
	}
}
