package routes

import (
	"github.com/gin-gonic/gin"
	"intelliagric-backend/internal/handlers"
	"intelliagric-backend/internal/repositories"
	"intelliagric-backend/internal/services"
)

func RegisterNewsRoutes(router *gin.RouterGroup, newsHander *handlers.NewsHandler) {
	// Initialize repository, service, and handler
	newsRepo := repositories.InitNewsRepository()
	newsService := services.InitNewsService(newsRepo)
	newsHandler := handlers.InitNewsHandler(newsService)

	{
		router.GET("/agriculture_news", newsHandler.GetAgricultureNews)
	}
}
