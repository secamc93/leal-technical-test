package dtos

type StoreResponse struct {
	ID               uint    `json:"id"`
	Name             string  `json:"name"`
	ConversionFactor float64 `json:"conversion_factor"`
}

type StoreRequest struct {
	Name             string  `json:"name"`
	ConversionFactor float64 `json:"conversion_factor"`
}
