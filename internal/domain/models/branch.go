package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	StoreID      uint          `json:"store_id" gorm:"not null"`
	Name         string        `json:"name" gorm:"type:varchar(100);not null"`
	Address      string        `json:"address" gorm:"type:varchar(200)"`
	Store        Store         `json:"store" gorm:"foreignKey:StoreID"`         // Relation to Store
	Campaigns    []Campaign    `json:"campaigns" gorm:"foreignKey:BranchID"`    // Relation to campaigns
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:BranchID"` // Relation to transactions
}
