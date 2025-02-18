package middleware

import (
	"net/http"

	"intelliagric-backend/internal/auth"
	"intelliagric-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("session_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "Unauthorized"})
			ctx.Abort()
			return
		}

		user, err := auth.ValidateJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{Error: "Invalid session"})
			ctx.Abort()
			return
		}

		// Store user ID in context for later use
		ctx.Set("userName", user.Username)
		ctx.Set("userEmail", user.Email)
		ctx.Next()
	}
}