package services

import (
	"fmt"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/dtos"
	"leal-technical-test/internal/infra/repository"

	"gorm.io/gorm"
)

// AccumulatedRewardService interface
type AccumulatedRewardService interface {
	GetAllRewards() ([]models.AccumulatedReward, error)
	GetRewardById(id uint) (*models.AccumulatedReward, error)
	GetRewardByUserAndStore(userID uint, storeID uint) (*models.AccumulatedReward, error)
	CreateReward(id uint, transaction *models.Transaction) error
	ClaimReward(claim dtos.ClaimRewardRequest) (string, error)
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

// CreateReward creates a new accumulated reward
func (s *accumulatedRewardService) CreateReward(storeId uint, transaction *models.Transaction) error {
	acumulatedReward := models.AccumulatedReward{
		UserID:              transaction.UserID,
		StoreID:             storeId,
		PointsAccumulated:   transaction.PointsEarned,
		CashbackAccumulated: transaction.CashbackEarned,
	}
	points, err := s.repo.GetByUserAndStore(transaction.UserID, storeId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// No record found, create a new one
			err := s.repo.Create(&acumulatedReward)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Record found, update it
		acumulatedReward.PointsAccumulated += points.PointsAccumulated
		acumulatedReward.CashbackAccumulated += points.CashbackAccumulated
		err = s.repo.UpdateAcumulateReward(transaction.UserID, &acumulatedReward)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *accumulatedRewardService) ClaimReward(claim dtos.ClaimRewardRequest) (string, error) {
	if claim.PointsAccumulated < claim.RewardRequired {
		return "", fmt.Errorf("insufficient points")
	}
	acumulate := models.AccumulatedReward{
		PointsAccumulated: claim.PointsAccumulated - claim.RewardRequired,
	}
	err := s.repo.UpdateAcumulateReward(claim.UserID, &acumulate)
	if err != nil {
		return "", err
	}
	return claim.Description, nil
}
