package handlers

import (
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/services"
	"intelliagric-backend/internal/utils"
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
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{Error: "Failed to get users"})
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
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}
	if err := handler.Service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{Error: "Failed to create user"})
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
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{Error: "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

