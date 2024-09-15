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
	CreateTransaction(transaction *models.Transaction) (float64, error)
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
func (s *transactionService) CreateTransaction(transaction *models.Transaction) (float64, error) {
	// Buscar sucursal
	branch, err := s.repoBranch.GetById(transaction.BranchID)
	if err != nil {
		return 0, fmt.Errorf("branch not found")
	}

	// Obtener factor de conversión y configurar fecha de la transacción
	conversionFactor := branch.Store.ConversionFactor
	transaction.Date = time.Now()
	transaction.RewardType = "points"

	// Buscar campaña activa en la sucursal y fecha específica
	campaign, _ := s.repoCampaign.FindByBranchAndDate(transaction.BranchID, transaction.Date)

	var points float64
	// Validación para la campaña de sucursal 1 (doble de puntos entre el 15 y 30 de mayo)
	if campaign != nil && campaign.Type == "double" {
		points = (transaction.Amount * conversionFactor) * 2
	} else if campaign != nil && campaign.Type == "additional" && transaction.Amount > 20000 {
		// Validación para la campaña de sucursal 2 (30% puntos adicionales para compras > 20000 entre el 15 y 20 de mayo)
		points = (transaction.Amount * conversionFactor) * 1.30
	} else {
		// Cálculo estándar si no aplica ninguna campaña
		points = transaction.Amount * conversionFactor
	}

	// Asignar los puntos ganados
	transaction.PointsEarned = points

	// Guardar la transacción
	err = s.repo.Create(transaction)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %v", err)
	}

	return points, nil
}
