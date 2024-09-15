package repository

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
)

// RewardRepository interface
type RewardRepository interface {
	GetAll() ([]models.Reward, error)
	GetById(id uint) (*models.Reward, error)
	GetByStoreId(storeID uint) ([]models.Reward, error)
	Delete(id uint) error
	Put(id uint, reward *models.Reward) error
	Create(reward *models.Reward) error
	Validate(description string) bool
}

// rewardRepository struct
type rewardRepository struct {
	db  config.IDatabaseConnection
	log config.ILogger
}

// NewRewardRepository constructor
func NewRewardRepository(db config.IDatabaseConnection) RewardRepository {
	log := config.NewLogger()
	return &rewardRepository{
		db:  db,
		log: log,
	}
}

// GetAll retrieves all rewards
func (r *rewardRepository) GetAll() ([]models.Reward, error) {
	var rewards []models.Reward
	if err := r.db.GetDB().
		Preload("Store").
		Find(&rewards).Error; err != nil {
		return nil, err
	}
	return rewards, nil
}

// GetById retrieves a reward by its ID
func (r *rewardRepository) GetById(id uint) (*models.Reward, error) {
	var reward models.Reward
	if err := r.db.GetDB().
		Preload("Store").
		First(&reward, id).Error; err != nil {
		return nil, err
	}
	return &reward, nil
}

// GetByStoreId retrieves rewards by StoreID
func (r *rewardRepository) GetByStoreId(storeID uint) ([]models.Reward, error) {
	var rewards []models.Reward
	if err := r.db.GetDB().
		Preload("Store").
		Where("store_id = ?", storeID).Find(&rewards).Error; err != nil {
		return nil, err
	}
	return rewards, nil
}

// Delete deletes a reward by its ID
func (r *rewardRepository) Delete(id uint) error {
	result := r.db.GetDB().Delete(&models.Reward{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("reward with ID %d not found", id)
	}
	return nil
}

// Update updates an existing reward
func (r *rewardRepository) Put(id uint, reward *models.Reward) error {
	result := r.db.GetDB().Model(&models.Reward{}).Where("id = ?", id).Updates(reward)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		r.log.Error("store not found")
		return fmt.Errorf("store not found")
	}
	return nil
}

// Create creates a new reward
func (r *rewardRepository) Create(reward *models.Reward) error {
	if err := r.db.GetDB().Create(reward).Error; err != nil {
		return err
	}
	return nil
}

// Validate exist description
func (r *rewardRepository) Validate(description string) bool {
	var reward models.Reward
	if err := r.db.GetDB().Where("description = ?", description).First(&reward).Error; err != nil {
		return false
	}
	return true
}
