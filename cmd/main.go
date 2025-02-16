package main

import (
	"log"
	"intelliagric-backend/config"
	"intelliagric-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv(config.DefaultLoadEnv)

	// initialize database
	db := &config.Database{}

	db.Connect()

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, db)

	// Start the server
	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
