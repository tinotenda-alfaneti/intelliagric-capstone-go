package routes

import (
	"github.com/gin-gonic/gin"
	"intelliagric-backend/internal/handlers"
)

func RegisterNewsRoutes(router *gin.RouterGroup, newsHandler *handlers.NewsHandler) {

	router.GET("/agriculture_news", newsHandler.GetAgricultureNews)

}