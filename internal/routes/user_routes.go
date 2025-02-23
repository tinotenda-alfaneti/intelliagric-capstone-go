package routes

import (
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/users", userHandler.GetUsers)
	protected.POST("/users", userHandler.CreateUser)
	protected.GET("/users/:id", userHandler.GetUserByID)
}
