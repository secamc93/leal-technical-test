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

// CampaignController struct
type CampaignController struct {
	service services.CampaignService
}

// NewCampaignController constructor
func NewCampaignController() *CampaignController {
	db := config.NewPostgresConnection()
	repo := repository.NewCampaignRepository(db)
	service := services.NewCampaignService(repo)

	return &CampaignController{
		service: service,
	}
}

// GetAllCampaigns handles GET requests to retrieve all campaigns
// @Summary Get all campaigns
// @Description Get all campaigns
// @Tags campaigns
// @Accept  json
// @Produce  json
// @Router /leal-test/campaigns [get]
func (c *CampaignController) GetAllCampaigns(ctx *gin.Context) {
	campaigns, err := c.service.GetAllCampaigns()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	campaignsDTO := adapters.ToCampaignDTOs(campaigns)

	ctx.JSON(http.StatusOK, campaignsDTO)
}

// GetCampaignById handles GET requests to retrieve a campaign by its ID
// @Summary Get campaign by ID
// @Description Get campaign by ID
// @Tags campaigns
// @Accept  json
// @Produce  json
// @Param id path int true "Campaign ID"
// @Router /leal-test/campaigns/{id} [get]
func (c *CampaignController) GetCampaignById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	campaign, err := c.service.GetCampaignById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	campaignDTO := adapters.ToCampaignDTO(campaign)
	ctx.JSON(http.StatusOK, campaignDTO)
}

// CreateCampaign handles POST requests to create a new campaign
// @Summary Create a new campaign
// @Description Create a new campaign
// @Tags campaigns
// @Accept  json
// @Produce  json
// @Param campaign body dtos.CampaignRequest true "Campaign data"
// @Router /leal-test/campaigns [post]
func (c *CampaignController) CreateCampaign(ctx *gin.Context) {
	var campaignDTO dtos.CampaignRequest
	if err := ctx.ShouldBindJSON(&campaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign := adapters.ToCampaignModel(campaignDTO)

	err := c.service.CreateCampaign(&campaign)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Campaign created successfully"})
}

// UpdateCampaign handles PUT requests to update a campaign
// @Summary Update a campaign
// @Description Update a campaign
// @Tags campaigns
// @Accept  json
// @Produce  json
// @Param id path int true "Campaign ID"
// @Param campaign body dtos.CampaignRequest true "Campaign data"
// @Router /leal-test/campaigns/{id} [put]
func (c *CampaignController) UpdateCampaign(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	var campaignDTO dtos.CampaignRequest
	if err := ctx.ShouldBindJSON(&campaignDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	campaign := adapters.ToCampaignModel(campaignDTO)

	err = c.service.UpdateCampaign(uint(id), &campaign)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Campaign updated successfully"})
}

// DeleteCampaign handles DELETE requests to delete a campaign by ID
// @Summary Delete campaign by ID
// @Description Delete a campaign by ID
// @Tags campaigns
// @Param id path int true "Campaign ID"
// @Router /leal-test/campaigns/{id} [delete]
func (c *CampaignController) DeleteCampaign(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	err = c.service.DeleteCampaign(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
