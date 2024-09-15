package dtos

import "time"

type CampaignResponse struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	BranchID   uint      `json:"branch_id"`
	Branch     string    `json:"branch"`
	Type       string    `json:"type"`
	Percentage float64   `json:"percentage"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date" `
}

type CampaignRequest struct {
	Name       string    `json:"name"`
	BranchID   uint      `json:"branch_id"`
	Type       string    `json:"type"`
	Percentage float64   `json:"percentage"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date" `
}
