package main

import (
	"log"
	"intelliagric-backend/config"
	"intelliagric-backend/internal/routes"
	_ "intelliagric-backend/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title IntelliAgric API
// @version 1.0
// @description This is the API documentation for IntelliAgric.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.intelliagric.com/support
// @contact.email support@intelliagric.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

func main() {

	config.LoadEnv(config.DefaultLoadEnv)

	// initialize database
	db := &config.Database{}

	db.Connect()

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, db)

	// Add Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
