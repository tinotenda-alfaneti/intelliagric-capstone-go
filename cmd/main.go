package main

import (
	"context"
	"intelliagric-backend/config"
	_ "intelliagric-backend/docs"
	"intelliagric-backend/internal/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func initTracer() func(context.Context) error {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	// For the demonstration, use AlwaysSample to sample every trace.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp.Shutdown
}
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

	// Initialize tracer
	shutdown := initTracer()
	defer func() {
		// Allow time for any remaining spans to be exported.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			log.Fatalf("Error shutting down tracer provider: %v", err)
		}
	}()

	// initialize database
	db := &config.Database{}

	db.Connect()

	router := gin.Default()

		// Add OpenTelemetry middleware for Gin.
	// This will automatically trace incoming HTTP requests.
	router.Use(otelgin.Middleware("intelliagric-backend"))

	// Register routes
	routes.RegisterRoutes(router, db)

	// Add Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	port := config.GetPort()
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
