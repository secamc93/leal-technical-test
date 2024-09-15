package models

import "gorm.io/gorm"

type AccumulatedReward struct {
	gorm.Model
	UserID              uint    `json:"user_id" gorm:"not null"`
	StoreID             uint    `json:"store_id" gorm:"not null"`
	PointsAccumulated   float64 `json:"points_accumulated" gorm:"type:decimal(10,2);default:0"`
	CashbackAccumulated float64 `json:"cashback_accumulated" gorm:"type:decimal(10,2);default:0"`
	User                User    `json:"user" gorm:"foreignKey:UserID"`
	Store               Store   `json:"store" gorm:"foreignKey:StoreID"`
}
