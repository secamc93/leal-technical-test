package repository

import (
	"errors"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"

	"gorm.io/gorm"
)

// BranchRepository interface
type BranchRepository interface {
	GetAll() ([]models.Branch, error)
	GetById(id uint) (*models.Branch, error)
	Delete(id uint) error
	Put(branch *models.Branch) error
	Post(branch *models.Branch) error
	ExistsByName(name string) bool
}

// branchRepository struct
type branchRepository struct {
	db config.IDatabaseConnection
}

// NewBranchRepository constructor
func NewBranchRepository(db config.IDatabaseConnection) BranchRepository {
	return &branchRepository{db: config.NewPostgresConnection()}
}

// GetAll retrieves all branches
func (r *branchRepository) GetAll() ([]models.Branch, error) {
	var branches []models.Branch
	if err := r.db.GetDB().
		Preload("Store").
		Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

// GetById retrieves a branch by its ID
func (r *branchRepository) GetById(id uint) (*models.Branch, error) {
	var branch models.Branch
	if err := r.db.GetDB().
		Preload("Store").
		First(&branch, id).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

// Delete deletes a branch by its ID
func (r *branchRepository) Delete(id uint) error {
	if err := r.db.GetDB().Delete(&models.Branch{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Put updates an existing branch
func (r *branchRepository) Put(branch *models.Branch) error {
	if err := r.db.GetDB().Save(branch).Error; err != nil {
		return err
	}
	return nil
}

// Post creates a new branch
func (r *branchRepository) Post(branch *models.Branch) error {
	if err := r.db.GetDB().Create(branch).Error; err != nil {
		return err
	}
	return nil
}

func (r *branchRepository) ExistsByName(name string) bool {
	var branch models.Branch
	if err := r.db.GetDB().Where("name = ?", name).First(&branch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
