package routes

import (
	"intelliagric-backend/config"
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/repositories"
	"intelliagric-backend/internal/services"
	"intelliagric-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *config.Database) {

	userRepo := repositories.InitUserRepository(db.DB)
	userService := services.InitUserService(userRepo)
	userHandler := handlers.InitUserHandler(userService)

	api := router.Group("/api")
	{
		api.POST("/signup", userHandler.SignUp)
        api.POST("/login", userHandler.Login)
        api.POST("/logout", userHandler.Logout)

        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware())
        protected.GET("/protected", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{"message": "You are authorized"})
        })
		protected.GET("/users", userHandler.GetUsers)
		protected.POST("/users", userHandler.CreateUser)
		protected.GET("/users/:id", userHandler.GetUserByID)
	}
}
