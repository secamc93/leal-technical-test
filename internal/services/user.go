package services

import (
	"fmt"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// UserService interface
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user *models.User) error
	CreateUser(user *models.User) error
}

// userService struct
type userService struct {
	repo repository.UserRepository
}

// NewUserService constructor
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// GetAllUsers retrieves all users
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// GetUserById retrieves a user by its ID
func (s *userService) GetUserById(id uint) (*models.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser removes a user by its ID
func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id uint, user *models.User) error {
	return s.repo.Update(id, user)
}

// CreateUser creates a new user
func (s *userService) CreateUser(user *models.User) error {
	exist := s.repo.GetByEmail(user.Email)
	if exist {
		return fmt.Errorf("user already exists")
	}
	return s.repo.Create(user)
}
