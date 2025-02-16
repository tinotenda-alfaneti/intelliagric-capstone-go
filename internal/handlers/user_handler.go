package handlers

import (
	"intelliagric-backend/internal/services"
	"intelliagric-backend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func InitUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}
