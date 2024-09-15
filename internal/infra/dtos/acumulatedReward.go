package dtos

type AccumulatedRewardResponse struct {
	Id                  uint    `json:"id"`
	UserID              uint    `json:"user_id"`
	User                string  `json:"user"`
	StoreID             uint    `json:"store_id"`
	Store               string  `json:"store"`
	PointsAccumulated   float64 `json:"points_accumulated"`
	CashbackAccumulated float64 `json:"cashback_accumulated"`
}
type AccumulatedRewardRequest struct {
	UserID              uint    `json:"user_id"`
	StoreID             uint    `json:"store_id"`
	PointsAccumulated   float64 `json:"points_accumulated"`
	CashbackAccumulated float64 `json:"cashback_accumulated"`
}
