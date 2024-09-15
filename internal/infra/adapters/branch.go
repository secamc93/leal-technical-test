package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToBranchDTO(branch *models.Branch) dtos.BranchResponse {
	if branch == nil {
		return dtos.BranchResponse{}
	}
	return dtos.BranchResponse{
		Id:               branch.ID,
		Name:             branch.Name,
		Address:          branch.Address,
		StoreID:          branch.StoreID,
		Store:            branch.Store.Name,
		ConversionFactor: branch.Store.ConversionFactor,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToBranchDTOs(branches []models.Branch) []dtos.BranchResponse {
	branchesDTO := make([]dtos.BranchResponse, len(branches))
	for i, branch := range branches {
		branchesDTO[i] = dtos.BranchResponse{
			Id:               branch.ID,
			Name:             branch.Name,
			Address:          branch.Address,
			StoreID:          branch.StoreID,
			Store:            branch.Store.Name,
			ConversionFactor: branch.Store.ConversionFactor,
		}
	}
	return branchesDTO
}

// Convierte un DTO en un modelo de dominio
func ToBranchModel(branch dtos.BranchRequest) models.Branch {
	return models.Branch{
		StoreID: branch.StoreID,
		Name:    branch.Name,
		Address: branch.Address,
	}
}
