package models

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	StoreID        int     `json:"store_id" gorm:"not null"`
	Description    string  `json:"description" gorm:"type:varchar(100)"`
	PointsRequired float64 `json:"points_required" gorm:"type:decimal(10,2)"`
	Store          Store   `json:"store" gorm:"foreignKey:StoreID"` // Relation to Store
}
