package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToTransactionDTO(transaction *models.Transaction) dtos.TransactionResponse {
	if transaction == nil {
		return dtos.TransactionResponse{}
	}
	return dtos.TransactionResponse{
		UserID:         transaction.UserID,
		User:           transaction.User.Name,
		BranchID:       transaction.BranchID,
		Branch:         transaction.Branch.Name,
		Amount:         transaction.Amount,
		Date:           transaction.Date,
		RewardType:     transaction.RewardType,
		PointsEarned:   transaction.PointsEarned,
		CashbackEarned: transaction.CashbackEarned,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToTransactionDTOs(transactions []models.Transaction) []dtos.TransactionResponse {
	transactionsDTO := make([]dtos.TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		transactionsDTO[i] = dtos.TransactionResponse{
			UserID:         transaction.UserID,
			User:           transaction.User.Name,
			BranchID:       transaction.BranchID,
			Branch:         transaction.Branch.Name,
			Amount:         transaction.Amount,
			Date:           transaction.Date,
			RewardType:     transaction.RewardType,
			PointsEarned:   transaction.PointsEarned,
			CashbackEarned: transaction.CashbackEarned,
		}
	}
	return transactionsDTO
}

// Convierte un DTO en un modelo de dominio
func ToTransactionModel(transaction dtos.TransactionRequest) *models.Transaction {
	return &models.Transaction{
		UserID:   transaction.UserID,
		BranchID: transaction.BranchID,
		Amount:   transaction.Amount,
	}
}
