package repository

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
)

// StoreRepository interface
type StoreRepository interface {
	GetAll() ([]models.Store, error)
	GetById(id uint) (*models.Store, error)
	Delete(id uint) error
	Put(id uint, store *models.Store) error
	Post(store *models.Store) error
}

// storeRepository struct
type storeRepository struct {
	db  config.IDatabaseConnection
	log config.ILogger
}

// NewStoreRepository constructor
func NewStoreRepository(db config.IDatabaseConnection) StoreRepository {
	return &storeRepository{
		db:  db,
		log: config.NewLogger(),
	}
}

// GetAll retrieves all stores
func (r *storeRepository) GetAll() ([]models.Store, error) {
	var stores []models.Store
	if err := r.db.GetDB().Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

// GetById retrieves a store by its ID
func (r *storeRepository) GetById(id uint) (*models.Store, error) {
	var store models.Store
	if err := r.db.GetDB().First(&store, id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

// Delete removes a store by its ID
func (r *storeRepository) Delete(id uint) error {
	if err := r.db.GetDB().Delete(&models.Store{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Put updates an existing store
func (r *storeRepository) Put(id uint, store *models.Store) error {
	result := r.db.GetDB().Model(&models.Store{}).Where("id = ?", id).Updates(store)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		r.log.Error("store not found")
		return fmt.Errorf("store not found")
	}
	return nil
}

// Post creates a new store
func (r *storeRepository) Post(store *models.Store) error {
	if err := r.db.GetDB().Create(store).Error; err != nil {
		return err
	}
	return nil
}
