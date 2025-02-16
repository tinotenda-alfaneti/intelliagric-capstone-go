package repositories

import (
	"intelliagric-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	err := repo.DB.Create(user).Error
	return err
}

func (repo *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	return user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    err := repo.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}