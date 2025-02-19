package routes

import (
	"intelliagric-backend/config"
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/repositories"
	"intelliagric-backend/internal/services"
	"intelliagric-backend/internal/utils"
	"intelliagric-backend/internal/middlewares"

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

	rateLimiter := utils.InitRateLimiter(5, 10)
	
	// API Group
	api := router.Group("/api")
	api.Use(middleware.RateLimitMiddleware(rateLimiter))

	RegisterUserRoutes(api, userHandler)
	RegisterNewsRoutes(api, newsHandler)
	RegisterAuthRoutes(api, authHandler)
}
