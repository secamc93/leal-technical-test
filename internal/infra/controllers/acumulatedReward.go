package controllers

import (
	"net/http"
	"strconv"

	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/adapters"
	"leal-technical-test/internal/infra/repository"
	"leal-technical-test/internal/services"

	"github.com/gin-gonic/gin"
)

// AccumulatedRewardController struct
type AccumulatedRewardController struct {
	service services.AccumulatedRewardService
}

// NewAccumulatedRewardController constructor
func NewAccumulatedRewardController() *AccumulatedRewardController {
	db := config.NewPostgresConnection()
	repo := repository.NewAccumulatedRewardRepository(db)
	service := services.NewAccumulatedRewardService(repo)

	return &AccumulatedRewardController{
		service: service,
	}
}

// GetAllRewards handles GET requests to retrieve all accumulated rewards
// @Summary Get all accumulated rewards
// @Description Get all accumulated rewards
// @Tags accumulated_rewards
// @Accept  json
// @Produce  json
// @Router /leal-test/acumulaterewards [get]
func (c *AccumulatedRewardController) GetAllRewards(ctx *gin.Context) {
	rewards, err := c.service.GetAllRewards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rewardsDTOs := adapters.ToAccumulateRewardDTOs(rewards)

	ctx.JSON(http.StatusOK, rewardsDTOs)
}

// GetRewardById handles GET requests to retrieve an accumulated reward by its ID
// @Summary Get accumulated reward by ID
// @Description Get accumulated reward by ID
// @Tags accumulated_rewards
// @Accept  json
// @Produce  json
// @Param id path int true "Reward ID"
// @Router /leal-test/acumulaterewards/{id} [get]
func (c *AccumulatedRewardController) GetRewardById(ctx *gin.Context) {
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
	rewardDTO := adapters.ToAccumulateRewardDTO(reward)
	ctx.JSON(http.StatusOK, rewardDTO)
}

// GetRewardByUserAndStore handles GET requests to retrieve an accumulated reward by UserID and StoreID
// @Summary Get accumulated reward by UserID and StoreID
// @Description Get accumulated reward by UserID and StoreID
// @Tags accumulated_rewards
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID"
// @Param store_id path int true "Store ID"
// @Router /leal-test/acumulaterewards/user/{user_id}/store/{store_id} [get]
func (c *AccumulatedRewardController) GetRewardByUserAndStore(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	storeID, err := strconv.Atoi(ctx.Param("store_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	reward, err := c.service.GetRewardByUserAndStore(uint(userID), uint(storeID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reward)
}

// CreateReward handles POST requests to create a new accumulated reward
// @Summary Create a new accumulated reward
// @Description Create a new accumulated reward
// @Tags accumulated_rewards
// @Accept  json
// @Produce  json
// @Param reward body dtos.AccumulatedRewardRequest true "Reward data"
// @Router /leal-test/acumulaterewards [post]
func (c *AccumulatedRewardController) CreateReward(ctx *gin.Context) {
	var reward models.AccumulatedReward
	if err := ctx.ShouldBindJSON(&reward); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.CreateReward(&reward)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reward)
}

// UpdateReward handles PUT requests to update an accumulated reward
// @Summary Update an accumulated reward
// @Description Update an accumulated reward
// @Tags accumulated_rewards
// @Accept  json
// @Produce  json
// @Param id path int true "Reward ID"
// @Param reward body dtos.AccumulatedRewardRequest true "Reward data"
// @Router /leal-test/acumulaterewards/{id} [put]
func (c *AccumulatedRewardController) UpdateReward(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reward ID"})
		return
	}

	var reward models.AccumulatedReward
	if err := ctx.ShouldBindJSON(&reward); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reward.ID = uint(id)
	err = c.service.UpdateReward(&reward)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reward)
}

// DeleteReward handles DELETE requests to delete an accumulated reward by ID
// @Summary Delete accumulated reward by ID
// @Description Delete an accumulated reward by ID
// @Tags accumulated_rewards
// @Param id path int true "Reward ID"
// @Router /leal-test/acumulaterewards/{id} [delete]
func (c *AccumulatedRewardController) DeleteReward(ctx *gin.Context) {
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
