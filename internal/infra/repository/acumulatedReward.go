package repository

import (
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
)

// AccumulatedRewardRepository interface
type AccumulatedRewardRepository interface {
	GetAll() ([]models.AccumulatedReward, error)
	GetById(id uint) (*models.AccumulatedReward, error)
	GetByUserAndStore(userID uint, storeID uint) (*models.AccumulatedReward, error)
	Delete(id uint) error
	Update(reward *models.AccumulatedReward) error
	Create(reward *models.AccumulatedReward) error
}

// accumulatedRewardRepository struct
type accumulatedRewardRepository struct {
	db config.IDatabaseConnection
}

// NewAccumulatedRewardRepository constructor
func NewAccumulatedRewardRepository(db config.IDatabaseConnection) AccumulatedRewardRepository {
	return &accumulatedRewardRepository{
		db: db,
	}
}

// GetAll retrieves all accumulated rewards
func (r *accumulatedRewardRepository) GetAll() ([]models.AccumulatedReward, error) {
	var rewards []models.AccumulatedReward
	if err := r.db.GetDB().
		Preload("User").
		Preload("Store").
		Find(&rewards).Error; err != nil {
		return nil, err
	}
	return rewards, nil
}

// GetById retrieves an accumulated reward by its ID
func (r *accumulatedRewardRepository) GetById(id uint) (*models.AccumulatedReward, error) {
	var reward models.AccumulatedReward
	if err := r.db.GetDB().
		Preload("User").
		Preload("Store").
		First(&reward, id).Error; err != nil {
		return nil, err
	}
	return &reward, nil
}

// GetByUserAndStore retrieves an accumulated reward by UserID and StoreID
func (r *accumulatedRewardRepository) GetByUserAndStore(userID uint, storeID uint) (*models.AccumulatedReward, error) {
	var reward models.AccumulatedReward
	if err := r.db.GetDB().Where("user_id = ? AND store_id = ?", userID, storeID).First(&reward).Error; err != nil {
		return nil, err
	}
	return &reward, nil
}

// Delete deletes an accumulated reward by its ID
func (r *accumulatedRewardRepository) Delete(id uint) error {
	if err := r.db.GetDB().Delete(&models.AccumulatedReward{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Update updates an existing accumulated reward
func (r *accumulatedRewardRepository) Update(reward *models.AccumulatedReward) error {
	if err := r.db.GetDB().Save(reward).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new accumulated reward
func (r *accumulatedRewardRepository) Create(reward *models.AccumulatedReward) error {
	if err := r.db.GetDB().Create(reward).Error; err != nil {
		return err
	}
	return nil
}
