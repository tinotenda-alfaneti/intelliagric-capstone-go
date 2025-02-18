package repositories

import (
	"encoding/json"
	"fmt"
	"intelliagric-backend/internal/models"
	"os"

	"github.com/go-resty/resty/v2"
)

// NewsRepository defines the interface for fetching news
type NewsRepository interface {
	FetchAgricultureNews() (models.NewsResponse, error)
}

// newsRepository implements NewsRepository
type newsRepository struct{}

// InitNewsRepository initializes a new repository instance
func InitNewsRepository() NewsRepository {
	return &newsRepository{}
}

// FetchAgricultureNews fetches news from GNews API
func (repo *newsRepository) FetchAgricultureNews() (models.NewsResponse, error) {
	apiKey := os.Getenv("GNEWS_API_KEY")
	baseURL := "https://gnews.io/api/v4/search"
	query := "agriculture+africa"
	lang := "en"
	maxResults := 10

	url := fmt.Sprintf("%s?q=%s&lang=%s&max=%d&apikey=%s", baseURL, query, lang, maxResults, apiKey)

	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Println("Error fetching news:", err)
		return models.NewsResponse{}, err
	}

	var newsData models.NewsResponse
	if err := json.Unmarshal(resp.Body(), &newsData); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return models.NewsResponse{}, err
	}

	return newsData, nil
}
