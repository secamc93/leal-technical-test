package repository

import (
	"errors"
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (*models.User, error)
	Delete(id uint) error
	Update(id uint, user *models.User) error
	Create(user *models.User) error
	GetByEmail(email string) bool
}

// userRepository struct
type userRepository struct {
	db config.IDatabaseConnection
}

// NewUserRepository constructor
func NewUserRepository(db config.IDatabaseConnection) UserRepository {
	return &userRepository{db: db}
}

// GetAll retrieves all users
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.GetDB().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetById retrieves a user by its ID
func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.GetDB().First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.GetDB().Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}
	return nil
}

// Update updates an existing user
func (r *userRepository) Update(id uint, user *models.User) error {
	if err := r.db.GetDB().Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new user

func (r *userRepository) Create(user *models.User) error {
	if err := r.db.GetDB().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByEmail(email string) bool {
	var user models.User
	if err := r.db.GetDB().Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
