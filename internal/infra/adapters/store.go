package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToStoreDTO(store models.Store) dtos.StoreResponse {
	return dtos.StoreResponse{
		ID:               store.ID,
		Name:             store.Name,
		ConversionFactor: store.ConversionFactor,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToStoreDTOs(stores []models.Store) []dtos.StoreResponse {
	storesDTO := make([]dtos.StoreResponse, len(stores))
	for i, store := range stores {
		storesDTO[i] = dtos.StoreResponse{
			ID:               store.ID,
			Name:             store.Name,
			ConversionFactor: store.ConversionFactor,
			// Mapear otros campos específicos aquí
		}
	}
	return storesDTO
}

// Convierte un DTO en un modelo de dominio
func ToStoreModel(dto dtos.StoreRequest) models.Store {
	return models.Store{
		Name:             dto.Name,
		ConversionFactor: dto.ConversionFactor,
	}
}
