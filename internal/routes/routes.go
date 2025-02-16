package routes

import (
	"intelliagric-backend/config"
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/repositories"
	"intelliagric-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *config.Database) {

	userRepo := repositories.InitUserRepository(db.DB)
	userService := services.InitUserService(userRepo)
	userHandler := handlers.InitUserHandler(userService)

	// Define routes
	api := router.Group("/api")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
	}
}
