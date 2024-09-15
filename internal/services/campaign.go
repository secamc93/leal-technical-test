package services

import (
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// CampaignService interface
type CampaignService interface {
	GetAllCampaigns() ([]models.Campaign, error)
	GetCampaignById(id uint) (*models.Campaign, error)
	DeleteCampaign(id uint) error
	UpdateCampaign(id uint, campaign *models.Campaign) error
	CreateCampaign(campaign *models.Campaign) error
}

// campaignService struct
type campaignService struct {
	repo repository.CampaignRepository
}

// NewCampaignService constructor
func NewCampaignService(repo repository.CampaignRepository) CampaignService {
	return &campaignService{
		repo: repo,
	}
}

// GetAllCampaigns retrieves all campaigns
func (s *campaignService) GetAllCampaigns() ([]models.Campaign, error) {
	campaigns, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

// GetCampaignById retrieves a campaign by its ID
func (s *campaignService) GetCampaignById(id uint) (*models.Campaign, error) {
	campaign, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

// DeleteCampaign deletes a campaign by its ID
func (s *campaignService) DeleteCampaign(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCampaign updates an existing campaign
func (s *campaignService) UpdateCampaign(id uint, campaign *models.Campaign) error {
	err := s.repo.Update(id, campaign)
	if err != nil {
		return err
	}
	return nil
}

// CreateCampaign creates a new campaign
func (s *campaignService) CreateCampaign(campaign *models.Campaign) error {
	err := s.repo.Create(campaign)
	if err != nil {
		return err
	}
	return nil
}
