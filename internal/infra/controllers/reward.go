package controllers

import (
	"net/http"
	"strconv"

	"leal-technical-test/config"
	"leal-technical-test/internal/infra/adapters"
	"leal-technical-test/internal/infra/dtos"
	"leal-technical-test/internal/infra/repository"
	"leal-technical-test/internal/services"

	"github.com/gin-gonic/gin"
)

// RewardController struct
type RewardController struct {
	service services.RewardService
}

// NewRewardController constructor
func NewRewardController() *RewardController {
	db := config.NewPostgresConnection()
	repo := repository.NewRewardRepository(db)
	service := services.NewRewardService(repo)

	return &RewardController{
		service: service,
	}
}

// GetAllRewards handles GET requests to retrieve all rewards
// @Summary Get all rewards
// @Description Get all rewards
// @Tags rewards
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /leal-test/rewards [get]
func (c *RewardController) GetAllRewards(ctx *gin.Context) {
	rewards, err := c.service.GetAllRewards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rewardsDTOs := adapters.ToRewardsDTOs(rewards)
	ctx.JSON(http.StatusOK, rewardsDTOs)
}

// GetRewardById handles GET requests to retrieve a reward by its ID
// @Summary Get reward by ID
// @Description Get reward by ID
// @Tags rewards
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Reward ID"
// @Router /leal-test/rewards/{id} [get]
func (c *RewardController) GetRewardById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reward ID"})
		return
	}

	reward, err := c.service.GetRewardById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rewardDTO := adapters.ToRewardsDTO(reward)
	ctx.JSON(http.StatusOK, rewardDTO)
}

// GetRewardsByStoreId handles GET requests to retrieve rewards by StoreID
// @Summary Get rewards by StoreID
// @Description Get rewards by StoreID
// @Tags rewards
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param store_id path int true "Store ID"
// @Router /leal-test/rewards/store/{store_id} [get]
func (c *RewardController) GetRewardsByStoreId(ctx *gin.Context) {
	storeID, err := strconv.Atoi(ctx.Param("store_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	rewards, err := c.service.GetRewardsByStoreId(uint(storeID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rewardsDTO := adapters.ToRewardsDTOs(rewards)
	ctx.JSON(http.StatusOK, rewardsDTO)
}

// CreateReward handles POST requests to create a new reward
// @Summary Create a new reward
// @Description Create a new reward
// @Tags rewards
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param reward body dtos.RewardRequest true "Reward data"
// @Router /leal-test/rewards [post]
func (c *RewardController) CreateReward(ctx *gin.Context) {
	var rewardDTO dtos.RewardRequest
	if err := ctx.ShouldBindJSON(&rewardDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reward := adapters.ToRewardModel(rewardDTO)
	err := c.service.CreateReward(&reward)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reward created successfully"})
}

// UpdateReward handles PUT requests to update a reward
// @Summary Update a reward
// @Description Update a reward
// @Tags rewards
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Reward ID"
// @Param reward body dtos.RewardRequest true "Reward data"
// @Router /leal-test/rewards/{id} [put]
func (c *RewardController) UpdateReward(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reward ID"})
		return
	}
	var rewardDTO dtos.RewardRequest
	if err := ctx.ShouldBindJSON(&rewardDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reward := adapters.ToRewardModel(rewardDTO)
	err = c.service.UpdateReward(uint(id), &reward)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reward updated successfully"})
}

// DeleteReward handles DELETE requests to delete a reward by ID
// @Summary Delete reward by ID
// @Description Delete a reward by ID
// @Tags rewards
// @Security ApiKeyAuth
// @Param id path int true "Reward ID"
// @Router /leal-test/rewards/{id} [delete]
func (c *RewardController) DeleteReward(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reward ID"})
		return
	}

	err = c.service.DeleteReward(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reward deleted successfully"})
}
