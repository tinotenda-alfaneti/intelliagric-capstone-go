package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"intelliagric-backend/internal/repositories"
)

// MockRepository mocks the NewsRepository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FetchAgricultureNews() (*repositories.NewsResponse, error) {
	args := m.Called()
	return args.Get(0).(*repositories.NewsResponse), args.Error(1)
}

func TestGetAgricultureNews(t *testing.T) {
	mockRepo := new(MockRepository)
	mockService := InitNewsService(mockRepo)

	// Fake response
	mockNews := &repositories.NewsResponse{
		Articles: []repositories.NewsArticle{
			{Title: "Fake News", Description: "This is a fake news article"},
		},
	}

	mockRepo.On("FetchAgricultureNews").Return(mockNews, nil)

	// Call service method
	news, err := mockService.GetAgricultureNews()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, news)
	assert.Equal(t, "Fake News", news.Articles[0].Title)
}
