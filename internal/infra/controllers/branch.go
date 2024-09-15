package controllers

import (
	"net/http"
	"strconv"

	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/adapters"
	"leal-technical-test/internal/infra/dtos"
	"leal-technical-test/internal/infra/repository"
	"leal-technical-test/internal/services"

	"github.com/gin-gonic/gin"
)

// BranchController struct
type BranchController struct {
	service services.BranchService
}

// NewBranchController constructor
func NewBranchController() *BranchController {
	db := config.NewPostgresConnection()
	repo := repository.NewBranchRepository(db)
	services := services.NewBranchService(repo)

	return &BranchController{
		service: services,
	}
}

// GetAllBranches handles GET requests to retrieve all branches
// @Summary Get all branches
// @Description Get all branches
// @Tags branches
// @Accept  json
// @Produce  json
// @Router /leal-test/branches [get]
func (c *BranchController) GetAllBranches(ctx *gin.Context) {
	branches, err := c.service.GetAllBranches()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	branchesDTO := adapters.ToBranchDTOs(branches)
	ctx.JSON(http.StatusOK, branchesDTO)
}

// GetBranchById handles GET requests to retrieve a branch by its ID
// @Summary Get branch by ID
// @Description Get branch by ID
// @Tags branches
// @Accept  json
// @Produce  json
// @Param id path int true "Branch ID"
// @Router /leal-test/branches/{id} [get]
func (c *BranchController) GetBranchById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}
	branch, err := c.service.GetBranchById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	branchDTO := adapters.ToBranchDTO(branch)
	ctx.JSON(http.StatusOK, branchDTO)
}

// CreateBranch handles POST requests to create a new branch
// @Summary Create a new branch
// @Description Create a new branch
// @Tags branches
// @Accept  json
// @Produce  json
// @Param branch body dtos.BranchRequest true "Branch data"
// @Router /leal-test/branches [post]
func (c *BranchController) CreateBranch(ctx *gin.Context) {
	var branchDTO dtos.BranchRequest
	if err := ctx.ShouldBindJSON(&branchDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	branch := adapters.ToBranchModel(branchDTO)
	err := c.service.CreateBranch(&branch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Branch created successfully"})
}

// UpdateBranch handles PUT requests to update a branch
// @Summary Update a branch
// @Description Update a branch
// @Tags branches
// @Accept  json
// @Produce  json
// @Param id path int true "Branch ID"
// @Param branch body dtos.BranchRequest true "Branch data"
// @Router /leal-test/branches/{id} [put]
func (c *BranchController) UpdateBranch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}

	var branch models.Branch
	if err := ctx.ShouldBindJSON(&branch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branch.ID = uint(id)
	err = c.service.UpdateBranch(&branch)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, branch)
}

// DeleteBranch handles DELETE requests to delete a branch by ID
// @Summary Delete branch by ID
// @Description Delete a branch by ID
// @Tags branches
// @Param id path int true "Branch ID"
// @Router /leal-test/branches/{id} [delete]
func (c *BranchController) DeleteBranch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch ID"})
		return
	}

	err = c.service.DeleteBranch(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Branch deleted successfully"})
}
