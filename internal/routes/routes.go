package routes

import (
	"intelliagric-backend/config"
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/repositories"
	"intelliagric-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes all routes
func RegisterRoutes(router *gin.Engine, db *config.Database) {
	// Initialize repositories
	userRepo := repositories.InitUserRepository(db.DB)
	newsRepo := repositories.InitNewsRepository()

	// Initialize services
	userService := services.InitUserService(userRepo)
	newsService := services.InitNewsService(newsRepo)
	authService := services.InitAuthService(userRepo)


	// Initialize handlers
	userHandler := handlers.InitUserHandler(userService)
	newsHandler := handlers.InitNewsHandler(newsService)
	authHandler := handlers.InitAuthHandler(authService)


	// API Group
	api := router.Group("/api")

	RegisterUserRoutes(api, userHandler)
	RegisterNewsRoutes(api, newsHandler)
	RegisterAuthRoutes(api, authHandler)
}
