package handlers

import (
	"fmt"
	"net/http"

	"intelliagric-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	Service services.NewsService
}

func InitNewsHandler(service services.NewsService) *NewsHandler {
	return &NewsHandler{Service: service}
}

// GetAgricultureNews handles GET /api/agriculture_news
func (handler *NewsHandler) GetAgricultureNews(ctx *gin.Context) {
	news, err := handler.Service.GetAgricultureNews()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to fetch news"})
		return
	}

	ctx.JSON(http.StatusOK, news)
}
