package services

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// UserService interface
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, user *models.User) error
	CreateUser(user *models.User) error
	Login(email string, password string) (string, error)
}

// userService struct
type userService struct {
	repo  repository.UserRepository
	token *config.TokenManager
}

// NewUserService constructor
func NewUserService(repo repository.UserRepository) UserService {
	token := config.NewTokenManager()
	return &userService{repo: repo, token: token}
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
	// Hash the user's password
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword

	// Save the user to the repository
	if err := s.repo.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (s *userService) HashPassword(password string) (string, error) {
	// Generate a hashed password with a default cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain text password
func (s *userService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *userService) Login(email string, password string) (string, error) {
	id, err := s.repo.GetIdByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	user, err := s.repo.GetById(id)
	if err != nil {
		return "", fmt.Errorf("error retrieving user")
	}

	validate := s.CheckPasswordHash(password, user.Password)
	if !validate {
		return "", fmt.Errorf("invalid password")
	}

	token, err := s.token.GenerateToken(user.Name)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}
