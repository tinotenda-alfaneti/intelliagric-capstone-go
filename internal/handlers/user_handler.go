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

type GetUsersResponse struct {
    Message string       `json:"message"`
    Users   []models.User `json:"users"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUsers retrieves all users
// @Summary Get all users
// @Description Fetches all users from the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.Service.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to get users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// CreateUser creates a new user
// @Summary Create a user
// @Description Creates a new user and stores it in the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

// GetUserByID retrieves a user by ID
// @Summary Get a user by ID
// @Description Fetches a user by ID from the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func (handler *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := handler.Service.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
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
func (handler *UserHandler) SignUp(ctx *gin.Context) {

	var req signUpReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request"})
		return
	}
	
	user :=  models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := handler.Service.SignUp(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
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
func (handler *UserHandler) Login(ctx *gin.Context) {

	var credentials loginReq

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid credentials"})
		return
	}

	token, err := handler.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Logout logs out a user
// @Summary Log out a user
// @Description Logs out a user
// @Tags Users
// @Success 200 {object} map[string]string
// @Router /logout [post]
func (handler *UserHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
