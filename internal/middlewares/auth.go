package middleware

import (
    "intelliagric-backend/internal/auth"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := ctx.GetHeader("Authorization")
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            ctx.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        _, err := auth.ValidateJWT(tokenString)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}
