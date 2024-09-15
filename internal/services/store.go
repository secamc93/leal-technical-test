package services

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// StoreService interface
type StoreService interface {
	GetAllStores() ([]models.Store, error)
	GetStoreById(id uint) (*models.Store, error)
	DeleteStore(id uint) error
	UpdateStore(id uint, store *models.Store) error
	CreateStore(store *models.Store) error
}

// storeService struct
type storeService struct {
	repo repository.StoreRepository
	log  config.ILogger
}

// NewStoreService constructor
func NewStoreService(repo repository.StoreRepository) StoreService {
	return &storeService{
		repo: repo,
		log:  config.NewLogger(),
	}
}

// GetAllStores retrieves all stores
func (s *storeService) GetAllStores() ([]models.Store, error) {
	stores, err := s.repo.GetAll()
	if err != nil {
		s.log.Error("Error retrieving stores: ", err)
		return nil, err
	}

	return stores, nil
}

// GetStoreById retrieves a store by its ID
func (s *storeService) GetStoreById(id uint) (*models.Store, error) {
	store, err := s.repo.GetById(id)
	if err != nil {
		s.log.Error("Error retrieving store by ID: ", err)
		return nil, err
	}
	return store, nil
}

// DeleteStore removes a store by its ID
func (s *storeService) DeleteStore(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.log.Error("Error deleting store: ", err)
		return err
	}
	return nil
}

// UpdateStore updates an existing store
func (s *storeService) UpdateStore(id uint, store *models.Store) error {
	// Verificar si la tienda existe
	existingStore, err := s.repo.GetById(id)
	if err != nil {
		s.log.Error("Error fetching store: ", err)
		return err
	}
	if existingStore == nil {
		s.log.Error("store not found")
		return fmt.Errorf("store not found")
	}

	// Actualizar la tienda
	err = s.repo.Put(id, store)
	if err != nil {
		s.log.Error("Error updating store: ", err)
		return err
	}
	return nil
}

// CreateStore creates a new store
func (s *storeService) CreateStore(store *models.Store) error {
	err := s.repo.Post(store)
	if err != nil {
		s.log.Error("Error creating store: ", err)
		return err
	}
	return nil
}
