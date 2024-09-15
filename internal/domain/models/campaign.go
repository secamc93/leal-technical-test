package models

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	Name       string    `json:"name" gorm:"type:varchar(100);unique"`
	BranchID   uint      `json:"branch_id" gorm:"not null"`
	Type       string    `json:"type" gorm:"type:varchar(20);not null;check:type IN ('double', 'additional')"`
	Percentage float64   `json:"percentage" gorm:"type:decimal(5,2)"`
	StartDate  time.Time `json:"start_date" gorm:"type:date;not null"`
	EndDate    time.Time `json:"end_date" gorm:"type:date;not null"`
	Branch     Branch    `json:"branch" gorm:"foreignKey:BranchID"` // Relation to Branch
}
