package dtos

import "time"

type TransactionResponse struct {
	UserID         uint      `json:"user_id"`
	User           string    `json:"user"`
	BranchID       uint      `json:"branch_id"`
	Branch         string    `json:"branch"`
	Amount         float64   `json:"amount" `
	Date           time.Time `json:"date" `
	RewardType     string    `json:"reward_type" `
	PointsEarned   float64   `json:"points_earned"`
	CashbackEarned float64   `json:"cashback_earned"`
}
type TransactionRequest struct {
	UserID   uint    `json:"user_id"`
	BranchID uint    `json:"branch_id"`
	Amount   float64 `json:"amount" `
}
