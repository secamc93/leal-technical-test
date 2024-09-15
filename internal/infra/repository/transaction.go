package repository

import (
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
)

// TransactionRepository interface
type TransactionRepository interface {
	GetAll() ([]models.Transaction, error)
	GetById(id uint) (*models.Transaction, error)
	GetByUserId(userID uint) ([]models.Transaction, error)
	Create(transaction *models.Transaction) error
}

// transactionRepository struct
type transactionRepository struct {
	db config.IDatabaseConnection
}

// NewTransactionRepository constructor
func NewTransactionRepository(db config.IDatabaseConnection) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

// GetAll retrieves all transactions
func (r *transactionRepository) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.GetDB().
		Preload("User").
		Preload("Branch").
		Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetById retrieves a transaction by its ID
func (r *transactionRepository) GetById(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.GetDB().
		Preload("User").
		Preload("Branch").
		First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetByUserId retrieves transactions by UserID
func (r *transactionRepository) GetByUserId(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.GetDB().
		Preload("User").
		Preload("Branch").
		Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// Create creates a new transaction
func (r *transactionRepository) Create(transaction *models.Transaction) error {
	if err := r.db.GetDB().Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
