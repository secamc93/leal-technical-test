package services

import (
	"fmt"
	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/repository"
)

// BranchService interface
type BranchService interface {
	GetAllBranches() ([]models.Branch, error)
	GetBranchById(id uint) (*models.Branch, error)
	DeleteBranch(id uint) error
	UpdateBranch(branch *models.Branch) error
	CreateBranch(branch *models.Branch) error
}

// branchService struct
type branchService struct {
	repo repository.BranchRepository
	log  config.ILogger
}

// NewBranchService constructor
func NewBranchService(repo repository.BranchRepository) BranchService {
	return &branchService{
		repo: repo,
		log:  config.NewLogger(),
	}
}

// GetAllBranches retrieves all branches
func (s *branchService) GetAllBranches() ([]models.Branch, error) {
	branches, err := s.repo.GetAll()
	if err != nil {
		s.log.Error("Error retrieving all branches: ", err)
		return nil, err
	}
	return branches, nil
}

// GetBranchById retrieves a branch by its ID
func (s *branchService) GetBranchById(id uint) (*models.Branch, error) {
	branch, err := s.repo.GetById(id)
	if err != nil {
		s.log.Error("Error retrieving branch by ID: ", err)
		return nil, err
	}
	return branch, nil
}

// DeleteBranch deletes a branch by its ID
func (s *branchService) DeleteBranch(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.log.Error("Error deleting branch by ID: ", err)
		return err
	}
	return nil
}

// UpdateBranch updates an existing branch
func (s *branchService) UpdateBranch(branch *models.Branch) error {
	exist := s.repo.ExistsByName(branch.Name)
	if !exist {
		return fmt.Errorf("branch does not exist")
	}

	err := s.repo.Put(branch)
	if err != nil {
		s.log.Error("Error updating branch: ", err)
		return err
	}
	return nil
}

// CreateBranch creates a new branch
func (s *branchService) CreateBranch(branch *models.Branch) error {
	exis := s.repo.ExistsByName(branch.Name)
	if exis {
		return fmt.Errorf("branch already exists")
	}

	err := s.repo.Post(branch)
	if err != nil {
		s.log.Error("Error creating branch: ", err)
		return err
	}
	return nil
}
