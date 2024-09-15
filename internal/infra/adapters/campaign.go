package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToCampaignDTO(campaign *models.Campaign) dtos.CampaignResponse {
	if campaign == nil {
		return dtos.CampaignResponse{}
	}
	return dtos.CampaignResponse{
		Id:         campaign.ID,
		Name:       campaign.Name,
		BranchID:   campaign.BranchID,
		Type:       campaign.Type,
		Percentage: campaign.Percentage,
		StartDate:  campaign.StartDate,
		EndDate:    campaign.EndDate,
		Branch:     campaign.Branch.Name,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToCampaignDTOs(campaigns []models.Campaign) []dtos.CampaignResponse {
	campaignsDTO := make([]dtos.CampaignResponse, len(campaigns))
	for i, campaign := range campaigns {
		campaignsDTO[i] = dtos.CampaignResponse{
			Id:         campaign.ID,
			Name:       campaign.Name,
			BranchID:   campaign.BranchID,
			Type:       campaign.Type,
			Percentage: campaign.Percentage,
			StartDate:  campaign.StartDate,
			EndDate:    campaign.EndDate,
			Branch:     campaign.Branch.Name,
		}
	}
	return campaignsDTO
}

// Convierte un DTO en un modelo de dominio
func ToCampaignModel(campaign dtos.CampaignRequest) models.Campaign {
	return models.Campaign{
		Name:       campaign.Name,
		BranchID:   campaign.BranchID,
		Type:       campaign.Type,
		Percentage: campaign.Percentage,
		StartDate:  campaign.StartDate,
		EndDate:    campaign.EndDate,
	}
}
