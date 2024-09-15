package repository

import (
	"errors"
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"

	"gorm.io/gorm"
)

// CampaignRepository interface
type CampaignRepository interface {
	GetAll() ([]models.Campaign, error)
	GetById(id uint) (*models.Campaign, error)
	Delete(id uint) error
	Update(id uint, campaign *models.Campaign) error
	Create(campaign *models.Campaign) error
}

// campaignRepository struct
type campaignRepository struct {
	db config.IDatabaseConnection
}

// NewCampaignRepository constructor
func NewCampaignRepository(db config.IDatabaseConnection) CampaignRepository {
	return &campaignRepository{
		db: db,
	}
}

// GetAll retrieves all campaigns
func (r *campaignRepository) GetAll() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	if err := r.db.GetDB().Find(&campaigns).Error; err != nil {
		return nil, err
	}
	return campaigns, nil
}

// GetById retrieves a campaign by its ID
func (r *campaignRepository) GetById(id uint) (*models.Campaign, error) {
	var campaign models.Campaign
	if err := r.db.GetDB().First(&campaign, id).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

// Delete deletes a campaign by its ID
func (r *campaignRepository) Delete(id uint) error {
	result := r.db.GetDB().Delete(&models.Campaign{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("campaign with ID %d not found", id)
	}
	return nil
}

// / Update updates an existing campaign
func (r *campaignRepository) Update(id uint, campaign *models.Campaign) error {
	// Verificar si la campaña existe
	var existingCampaign models.Campaign
	if err := r.db.GetDB().First(&existingCampaign, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("campaign with ID %d not found", id)
		}
		return err
	}

	// Establecer el ID de la campaña
	campaign.ID = id

	// Guardar la campaña actualizada
	if err := r.db.GetDB().Save(campaign).Error; err != nil {
		return err
	}
	return nil
}

// Create creates a new campaign
func (r *campaignRepository) Create(campaign *models.Campaign) error {
	if err := r.db.GetDB().Create(campaign).Error; err != nil {
		return err
	}
	return nil
}
