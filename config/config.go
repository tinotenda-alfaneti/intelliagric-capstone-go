package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)
type LoadEnvFunc func() error

// LoadEnv loads the environment variables
var LoadEnv = func(loadEnvFunc LoadEnvFunc) {
    err := loadEnvFunc()
    if err != nil {
        log.Println("No .env file found")
    }
}

// DefaultLoadEnv is the default function to load environment variables
func DefaultLoadEnv() error {
    return godotenv.Load()
}

// GetPort returns the server port
func GetPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    return port
}