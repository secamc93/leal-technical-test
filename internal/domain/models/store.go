package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name             string   `json:"name" gorm:"type:varchar(100);not null"`
	ConversionFactor float64  `json:"conversion_factor" gorm:"type:decimal(5,2);default:1.0"`
	Branches         []Branch `json:"branches" gorm:"foreignKey:StoreID"` // Relation to branches
	Rewards          []Reward `json:"rewards" gorm:"foreignKey:StoreID"`  // Relation to rewards
}
