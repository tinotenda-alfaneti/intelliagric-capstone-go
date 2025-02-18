package services

import (
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/repositories"

)

type UserService struct {
	Repo *repositories.UserRepository
}

func InitUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) GetUsers() ([]models.User, error) {
	return service.Repo.GetAllUsers()
}

func (service *UserService) CreateUser(user *models.User) error {
	// TODO: Add validation logic here
	return service.Repo.CreateUser(user)
}

func (service *UserService) GetUserByID(id string) (models.User, error) {
	return service.Repo.GetUserByID(id)
}