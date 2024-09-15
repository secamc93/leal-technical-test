package services

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// AccumulatedRewardService interface
type AccumulatedRewardService interface {
	GetAllRewards() ([]models.AccumulatedReward, error)
	GetRewardById(id uint) (*models.AccumulatedReward, error)
	GetRewardByUserAndStore(userID uint, storeID uint) (*models.AccumulatedReward, error)
	DeleteReward(id uint) error
	UpdateReward(reward *models.AccumulatedReward) error
	CreateReward(reward *models.AccumulatedReward) error
}

// accumulatedRewardService struct
type accumulatedRewardService struct {
	repo repository.AccumulatedRewardRepository
}

// NewAccumulatedRewardService constructor
func NewAccumulatedRewardService(repo repository.AccumulatedRewardRepository) AccumulatedRewardService {
	return &accumulatedRewardService{
		repo: repo,
	}
}

// GetAllRewards retrieves all accumulated rewards
func (s *accumulatedRewardService) GetAllRewards() ([]models.AccumulatedReward, error) {
	rewards, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

// GetRewardById retrieves an accumulated reward by its ID
func (s *accumulatedRewardService) GetRewardById(id uint) (*models.AccumulatedReward, error) {
	reward, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return reward, nil
}

// GetRewardByUserAndStore retrieves an accumulated reward by UserID and StoreID
func (s *accumulatedRewardService) GetRewardByUserAndStore(userID uint, storeID uint) (*models.AccumulatedReward, error) {
	reward, err := s.repo.GetByUserAndStore(userID, storeID)
	if err != nil {
		return nil, err
	}
	return reward, nil
}

// DeleteReward deletes an accumulated reward by its ID
func (s *accumulatedRewardService) DeleteReward(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateReward updates an existing accumulated reward
func (s *accumulatedRewardService) UpdateReward(reward *models.AccumulatedReward) error {
	err := s.repo.Update(reward)
	if err != nil {
		return err
	}
	return nil
}

// CreateReward creates a new accumulated reward
func (s *accumulatedRewardService) CreateReward(reward *models.AccumulatedReward) error {
	err := s.repo.Create(reward)
	if err != nil {
		return err
	}
	return nil
}
