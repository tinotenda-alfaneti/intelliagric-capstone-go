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

	api := router.Group("/api")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUserByID)
	}
}
