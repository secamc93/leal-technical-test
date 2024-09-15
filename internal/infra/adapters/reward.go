package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

// Convierte un modelo de dominio a un DTO
func ToRewardsDTO(reward *models.Reward) dtos.RewardResponse {
	if reward == nil {
		return dtos.RewardResponse{}
	}
	return dtos.RewardResponse{
		Id:             reward.ID,
		StoreID:        reward.StoreID,
		Store:          reward.Store.Name,
		Description:    reward.Description,
		PointsRequired: reward.PointsRequired,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToRewardsDTOs(rewards []models.Reward) []dtos.RewardResponse {
	rewardDTO := make([]dtos.RewardResponse, len(rewards))
	for i, reward := range rewards {
		rewardDTO[i] = dtos.RewardResponse{
			Id:             reward.ID,
			StoreID:        reward.StoreID,
			Store:          reward.Store.Name,
			Description:    reward.Description,
			PointsRequired: reward.PointsRequired,
		}
	}
	return rewardDTO
}

// Convierte un DTO en un modelo de dominio
func ToRewardModel(reward dtos.RewardRequest) models.Reward {
	return models.Reward{
		StoreID:        reward.StoreID,
		Description:    reward.Description,
		PointsRequired: reward.PointsRequired,
	}
}
