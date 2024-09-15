package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID         uint      `json:"user_id" gorm:"not null"`
	BranchID       uint      `json:"branch_id" gorm:"not null"`
	Amount         float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Date           time.Time `json:"date" gorm:"type:timestamp;default:current_timestamp"`
	RewardType     string    `json:"reward_type" gorm:"type:varchar(20);not null;check:reward_type IN ('points', 'cashback')"`
	PointsEarned   float64   `json:"points_earned" gorm:"type:decimal(10,2)"`
	CashbackEarned float64   `json:"cashback_earned" gorm:"type:decimal(10,2)"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`     // Relation to User
	Branch         Branch    `json:"branch" gorm:"foreignKey:BranchID"` // Relation to Branch
}
