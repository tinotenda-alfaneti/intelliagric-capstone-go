package services

import (
	"errors"
	"fmt"
	"intelliagric-backend/internal/auth"
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
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

func (service *UserService) SignUp(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Printf("Entered Password: [%s] (len=%d)\n", user.Password, len(user.Password))
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return service.Repo.CreateUser(user)
}

func (service *UserService) Login(email, password string) (string, error) {
    user, err := service.Repo.GetUserByEmail(email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Password mismatch:", err)
        return "", errors.New("invalid email or password")
    }

    token, err := auth.GenerateJWT(user.Name, user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}