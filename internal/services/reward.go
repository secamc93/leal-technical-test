package services

import (
	"fmt"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// RewardService interface
type RewardService interface {
	GetAllRewards() ([]models.Reward, error)
	GetRewardById(id uint) (*models.Reward, error)
	GetRewardsByStoreId(storeID uint) ([]models.Reward, error)
	DeleteReward(id uint) error
	UpdateReward(id uint, reward *models.Reward) error
	CreateReward(reward *models.Reward) error
}

// rewardService struct
type rewardService struct {
	repo repository.RewardRepository
}

// NewRewardService constructor
func NewRewardService(repo repository.RewardRepository) RewardService {
	return &rewardService{
		repo: repo,
	}
}

// GetAllRewards retrieves all rewards
func (s *rewardService) GetAllRewards() ([]models.Reward, error) {
	rewards, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

// GetRewardById retrieves a reward by its ID
func (s *rewardService) GetRewardById(id uint) (*models.Reward, error) {
	reward, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return reward, nil
}

// GetRewardsByStoreId retrieves rewards by StoreID
func (s *rewardService) GetRewardsByStoreId(storeID uint) ([]models.Reward, error) {
	rewards, err := s.repo.GetByStoreId(storeID)
	if err != nil {
		return nil, err
	}
	return rewards, nil
}

// DeleteReward deletes a reward by its ID
func (s *rewardService) DeleteReward(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateReward updates an existing reward
func (s *rewardService) UpdateReward(id uint, reward *models.Reward) error {
	existe := s.repo.Validate(reward.Description)
	if existe {
		return fmt.Errorf("reward already exists")
	}
	err := s.repo.Put(id, reward)
	if err != nil {
		return err
	}
	return nil
}

// CreateReward creates a new reward
func (s *rewardService) CreateReward(reward *models.Reward) error {
	exist := s.repo.Validate(reward.Description)
	if exist {
		return fmt.Errorf("reward already exists")
	}
	err := s.repo.Create(reward)
	if err != nil {
		return err
	}
	return nil
}
