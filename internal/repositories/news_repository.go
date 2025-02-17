package repositories

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// NewsArticle represents a single news article structure
type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
}

// NewsResponse represents the response from the API
type NewsResponse struct {
	Articles []NewsArticle `json:"articles"`
}

// NewsRepository defines the interface for fetching news
type NewsRepository interface {
	FetchAgricultureNews() (*NewsResponse, error)
}

// newsRepository implements NewsRepository
type newsRepository struct{}

// InitNewsRepository initializes a new repository instance
func InitNewsRepository() NewsRepository {
	return &newsRepository{}
}

// FetchAgricultureNews fetches news from GNews API
func (repo *newsRepository) FetchAgricultureNews() (*NewsResponse, error) {
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
		return nil, err
	}

	var newsData NewsResponse
	if err := json.Unmarshal(resp.Body(), &newsData); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}

	return &newsData, nil
}
