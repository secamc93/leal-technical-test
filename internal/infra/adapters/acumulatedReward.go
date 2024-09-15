package adapters

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
)

func ToAccumulateRewardDTO(reward *models.AccumulatedReward) dtos.AccumulatedRewardResponse {
	if reward == nil {
		return dtos.AccumulatedRewardResponse{}
	}
	return dtos.AccumulatedRewardResponse{
		Id:                  reward.ID,
		UserID:              reward.UserID,
		User:                reward.User.Name,
		StoreID:             reward.StoreID,
		Store:               reward.Store.Name,
		PointsAccumulated:   reward.PointsAccumulated,
		CashbackAccumulated: reward.CashbackAccumulated,
	}
}

// Convierte una lista de modelos de dominio a una lista de DTOs
func ToAccumulateRewardDTOs(rewards []models.AccumulatedReward) []dtos.AccumulatedRewardResponse {
	rewardsDTO := make([]dtos.AccumulatedRewardResponse, len(rewards))
	for i, reward := range rewards {
		rewardsDTO[i] = dtos.AccumulatedRewardResponse{
			Id:                  reward.ID,
			UserID:              reward.UserID,
			User:                reward.User.Name,
			StoreID:             reward.StoreID,
			Store:               reward.Store.Name,
			PointsAccumulated:   reward.PointsAccumulated,
			CashbackAccumulated: reward.CashbackAccumulated,
		}
	}
	return rewardsDTO
}

// Convierte un DTO en un modelo de dominio
func ToAccumulateRewardModel(reward dtos.AccumulatedRewardRequest) models.AccumulatedReward {
	return models.AccumulatedReward{
		UserID:              reward.UserID,
		StoreID:             reward.StoreID,
		PointsAccumulated:   reward.PointsAccumulated,
		CashbackAccumulated: reward.CashbackAccumulated,
	}
}
