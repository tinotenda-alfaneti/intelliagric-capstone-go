package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/utils"

	"github.com/sony/gobreaker"
)

// NewsRepository defines the interface for fetching news
type NewsRepository interface {
	FetchAgricultureNews() (models.NewsResponse, error)
}


// newsRepository implements NewsRepository.
type newsRepository struct {
	apiURL         string
	rateLimiter    *utils.RateLimiter
	circuitBreaker *gobreaker.CircuitBreaker
}

// InitNewsRepository initializes a new newsRepository instance with rate limiting and a circuit breaker.
func InitNewsRepository() NewsRepository {
	
	rateLimiter := utils.InitRateLimiter(5, 10)

	// Set up circuit breaker settings.
	cbSettings := gobreaker.Settings{
		Name:        "AgricultureNewsCB",
		Timeout:     5 * time.Second,
		MaxRequests: 3, // number of consecutive requests allowed when closed
		Interval:    10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip breaker if more than 50% of requests fail.
			return counts.ConsecutiveFailures > 3
		},
	}
	cicuitBreaker := gobreaker.NewCircuitBreaker(cbSettings)

	apiKey := os.Getenv("GNEWS_API_KEY")
	baseURL := "https://gnews.io/api/v4/search"
	query := "agriculture+africa"
	lang := "en"
	maxResults := 10

	apiUrl := fmt.Sprintf("%s?q=%s&lang=%s&max=%d&apikey=%s", baseURL, query, lang, maxResults, apiKey)

	return &newsRepository{
		apiURL:         apiUrl,
		rateLimiter:    rateLimiter,
		circuitBreaker: cicuitBreaker,
	}
}


// FetchAgricultureNews fetches news from GNews API
func (repo *newsRepository) FetchAgricultureNews() (models.NewsResponse, error) {

	repo.rateLimiter.Wait()

	// Wrap the HTTP call with the circuit breaker.
	result, err := repo.circuitBreaker.Execute(func() (interface{}, error) {
		resp, err := http.Get(repo.apiURL)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("failed to fetch agriculture news: " + resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var newsResponse models.NewsResponse
		if err := json.Unmarshal(body, &newsResponse); err != nil {
			return nil, err
		}


		return newsResponse, nil
	})

	if err != nil {
		return models.NewsResponse{}, err
	}

	newsData, ok := result.(models.NewsResponse)
	if !ok {
		return models.NewsResponse{}, errors.New("unexpected result type")
	}

	return newsData, nil

}
