package dtos

type BranchResponse struct {
	Id               uint    `json:"id"`
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	StoreID          uint    `json:"store_id"`
	Store            string  `json:"store"`
	ConversionFactor float64 `json:"connection_factor"`
}

type BranchRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	StoreID uint   `json:"store_id"`
}
