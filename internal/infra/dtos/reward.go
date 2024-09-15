package dtos

type RewardResponse struct {
	Id             uint    `json:"id"`
	StoreID        int     `json:"store_id"`
	Store          string  `json:"store"`
	Description    string  `json:"description"`
	PointsRequired float64 `json:"points_required" `
}

type RewardRequest struct {
	StoreID        int     `json:"store_id"`
	Description    string  `json:"description"`
	PointsRequired float64 `json:"points_required" `
}
