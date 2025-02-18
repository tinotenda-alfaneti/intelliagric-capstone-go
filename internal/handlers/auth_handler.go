package handlers

import (
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/utils"
	"net/http"
	
	"intelliagric-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func InitAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}


// SignUp creates a new user account
// @Summary Sign up a new user
// @Description Creates a new user account and stores it in the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Router /signup [post]
func (handler *AuthHandler) SignUp(ctx *gin.Context) {

	var req signUpReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid request"})
		return
	}
	
	user :=  models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := handler.Service.SignUp(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login authenticates a user and returns a token
// @Summary Log in a user
// @Description Authenticates a user and returns a token
// @Tags Users
// @Accept  json
// @Produce  json
// @Param credentials body loginReq true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func (handler *AuthHandler) Login(ctx *gin.Context) {

	var credentials loginReq

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	token, err := handler.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse{Error: err.Error()})
		return
	}

		// Set the JWT token in a secure HTTP-only cookie
		ctx.SetCookie("session_token", token, 
		3600,    // Expiry: 1 hour
		"/",     // Scope: entire domain
		"",      // No specific domain (uses the current one)
		false,    // Not Secure: also sent over HTTP
		true,    // HttpOnly: Not accessible via JavaScript
	)

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Logout logs out a user
// @Summary Log out a user
// @Description Logs out a user
// @Tags Users
// @Success 200 {object} map[string]string
// @Router /logout [post]
func (handler *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("session_token", "", 
		-1,  // Expire immediately
		"/",
		"",
		true,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}