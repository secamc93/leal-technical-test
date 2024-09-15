package services

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// TransactionService interface
type TransactionService interface {
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionById(id uint) (*models.Transaction, error)
	GetTransactionsByUserId(userID uint) ([]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
}

// transactionService struct
type transactionService struct {
	repo repository.TransactionRepository
}

// NewTransactionService constructor
func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

// GetAllTransactions retrieves all transactions
func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionById retrieves a transaction by its ID
func (s *transactionService) GetTransactionById(id uint) (*models.Transaction, error) {
	transaction, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetTransactionsByUserId retrieves transactions by UserID
func (s *transactionService) GetTransactionsByUserId(userID uint) ([]models.Transaction, error) {
	transactions, err := s.repo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// CreateTransaction creates a new transaction
func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	err := s.repo.Create(transaction)
	if err != nil {
		return err
	}
	return nil
}
