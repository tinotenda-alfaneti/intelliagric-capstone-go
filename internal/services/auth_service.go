package services

import (
	"errors"
	"fmt"
	"intelliagric-backend/internal/auth"
	"intelliagric-backend/internal/models"
	"intelliagric-backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repositories.UserRepository
}

func InitAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (service *AuthService) SignUp(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return service.Repo.CreateUser(user)
}

func (service *AuthService) Login(email, password string) (string, error) {
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