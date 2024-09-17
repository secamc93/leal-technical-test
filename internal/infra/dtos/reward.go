package dtos

type RewardResponse struct {
	Id             uint    `json:"id"`
	StoreID        uint    `json:"store_id"`
	Store          string  `json:"store"`
	Description    string  `json:"description"`
	PointsRequired float64 `json:"points_required" `
}

type RewardRequest struct {
	StoreID        uint    `json:"store_id"`
	Description    string  `json:"description"`
	PointsRequired float64 `json:"points_required" `
}

type ClaimRewardRequest struct {
	UserID            uint    `json:"user_id"`
	PointsAccumulated float64 `json:"points_accumulated"`
	RewardID          uint    `json:"reward_id"`
	RewardRequired    float64 `json:"reward_required"`
	StoreID           uint    `json:"store_id"`
	Description       string  `json:"description"`
}
