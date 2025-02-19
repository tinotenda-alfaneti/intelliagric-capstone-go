package middleware

import (
	"github.com/gin-gonic/gin"
	"intelliagric-backend/internal/utils"
)

// RateLimitMiddleware returns a Gin middleware function that waits for a token
func RateLimitMiddleware(rateLimiter *utils.RateLimiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rateLimiter.Wait()
		ctx.Next()
	}
}
