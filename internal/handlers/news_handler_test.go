package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"intelliagric-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNewsService mocks the NewsService
type MockNewsService struct {
	mock.Mock
}

func (m *MockNewsService) GetAgricultureNews() (models.NewsResponse, error) {
	args := m.Called()
	return args.Get(0).(models.NewsResponse), args.Error(1)
}

func TestGetAgricultureNewsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockNewsService)
	handler := InitNewsHandler(mockService)

	// Fake news response
	mockNews := models.NewsResponse{
		Articles: []models.NewsArticle{
			{Title: "Fake News", Description: "Testing handler"},
		},
	}
	mockService.On("GetAgricultureNews").Return(mockNews, nil)

	// Set up Gin router
	router := gin.Default()
	router.GET("/api/agriculture_news", handler.GetAgricultureNews)

	// Perform HTTP request
	req, _ := http.NewRequest("GET", "/api/agriculture_news", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Fake News")
}
