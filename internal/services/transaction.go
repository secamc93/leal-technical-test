package services

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
	"time"
)

// TransactionService interface
type TransactionService interface {
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionById(id uint) (*models.Transaction, error)
	GetTransactionsByUserId(userID uint) ([]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, uint, error)
}

// transactionService struct
type transactionService struct {
	log          config.ILogger
	repo         repository.TransactionRepository
	repoBranch   repository.BranchRepository
	repoCampaign repository.CampaignRepository
}

// NewTransactionService constructor
func NewTransactionService(
	repo repository.TransactionRepository,
	repoBranch repository.BranchRepository,
	repoCampaign repository.CampaignRepository,

) TransactionService {
	log := config.NewLogger()
	return &transactionService{
		log:          log,
		repo:         repo,
		repoBranch:   repoBranch,
		repoCampaign: repoCampaign,
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
func (s *transactionService) CreateTransaction(transaction *models.Transaction) (*models.Transaction, uint, error) {
	// Buscar sucursal
	branch, err := s.repoBranch.GetById(transaction.BranchID)
	if err != nil {
		return nil, 0, fmt.Errorf("branch not found")
	}

	campaign, err := s.repoCampaign.FindByBranchAndDate(transaction.BranchID, time.Now())
	if err != nil {
		transaction.PointsEarned = transaction.Amount * branch.Store.ConversionFactor
	} else {
		if campaign.Type == "double" {
			transaction.PointsEarned = (transaction.Amount * branch.Store.ConversionFactor) * 2
		} else if campaign.Type == "additional" && transaction.Amount > 20000 {
			transaction.PointsEarned = (transaction.Amount * branch.Store.ConversionFactor) * 1.30
		}
	}

	transaction.RewardType = "points"
	s.log.Info("transaction.PointsEarned: ", transaction.PointsEarned)
	err = s.repo.Create(transaction)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create transaction: %v", err)
	}

	fmt.Println("transactionService.CreateTransaction", branch.StoreID)
	return transaction, branch.StoreID, nil
}
