package models

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