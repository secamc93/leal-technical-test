package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string              `json:"name" gorm:"type:varchar(100);not null"`
	Email        string              `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone        string              `json:"phone" gorm:"type:varchar(20)"`
	Password     string              `json:"password" gorm:"type:varchar(255);not null"` // Password field
	Rewards      []AccumulatedReward `json:"rewards" gorm:"foreignKey:UserID"`           // Relation to accumulated rewards
	Transactions []Transaction       `json:"transactions" gorm:"foreignKey:UserID"`      // Relation to transactions
}
