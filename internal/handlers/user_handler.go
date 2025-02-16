package handlers

import (
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/services"
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

func (handler *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := handler.Service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}


func (handler *UserHandler) SignUp(ctx *gin.Context) {

	type signUpReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req signUpReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	
	user :=  models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := handler.Service.SignUp(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	type loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var credentials loginReq

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := handler.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// client should remove the token
func (handler *UserHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
