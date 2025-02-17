package services

import (
	"intelliagric-backend/internal/repositories"
)

// NewsService defines the interface for news-related operations
type NewsService interface {
	GetAgricultureNews() (*repositories.NewsResponse, error)
}

// newsService implements NewsService
type newsService struct {
	Repo repositories.NewsRepository
}

func InitNewsService(repo repositories.NewsRepository) NewsService {
	return &newsService{Repo: repo}
}

// GetAgricultureNews calls the repository to fetch news
func (service *newsService) GetAgricultureNews() (*repositories.NewsResponse, error) {
	return service.Repo.FetchAgricultureNews()
}
