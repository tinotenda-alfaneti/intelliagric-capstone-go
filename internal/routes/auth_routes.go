package routes

import (
	"github.com/gin-gonic/gin"
	"intelliagric-backend/internal/handlers"
)

func RegisterAuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {

	router.POST("/signup", authHandler.SignUp)
	router.POST("/login", authHandler.Login)
	router.POST("/logout", authHandler.Logout)

}