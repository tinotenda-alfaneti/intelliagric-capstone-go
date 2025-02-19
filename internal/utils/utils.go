package utils

import (
	"time"
)

type ErrorResponse struct {
    Error string `json:"error"`
}

// RateLimiter is a simple token bucket rate limiter.
type RateLimiter struct {
	tokens chan struct{}
}

// InitRateLimiter creates a new RateLimiter that refills tokens.
func InitRateLimiter(rate int, capacity int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, capacity),
	}

	// Pre-fill the bucket with tokens.
	for i := 0; i < capacity; i++ {
		rl.tokens <- struct{}{}
	}

	// This goroutine refills the bucket at the given rate.
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
				// Successfully added a token.
			default:
				// Bucket is full, skip adding a token.
			}
		}
	}()

	return rl
}

// Wait blocks until a token is available.
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}
