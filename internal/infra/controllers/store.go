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

// StoreController struct
type StoreController struct {
	service services.StoreService
}

// NewStoreController constructor
func NewStoreController() *StoreController {
	db := config.NewPostgresConnection()
	repo := repository.NewStoreRepository(db)
	services := services.NewStoreService(repo)

	return &StoreController{
		service: services,
	}
}

// GetAllStores handles GET requests to retrieve all stores

// GetAllStores godoc
// @Summary Get all stores
// @Description Get all stores
// @Tags stores
// @Accept  json
// @Produce  json
// @Router /leal-test/stores [get]
func (c *StoreController) GetAllStores(ctx *gin.Context) {
	stores, err := c.service.GetAllStores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	storesDTO := adapters.ToStoreDTOs(stores)
	ctx.JSON(http.StatusOK, storesDTO)
}

// GetStoreById handles GET requests to retrieve a store by its ID

// GetStoreById godoc
// @Summary Get store by ID
// @Description Get store by ID
// @Tags stores
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Router /leal-test/stores/{id} [get]
func (c *StoreController) GetStoreById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	store, err := c.service.GetStoreById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	storeDTO := adapters.ToStoreDTO(*store)

	ctx.JSON(http.StatusOK, storeDTO)
}

// DeleteStore handles DELETE requests to remove a store by its ID

// DeleteStore godoc
// @Summary Delete store by ID
// @Description Delete store by ID
// @Tags stores
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Router /leal-test/stores/{id} [delete]
func (c *StoreController) DeleteStore(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	err = c.service.DeleteStore(uint(id))
	if err != nil {
		if err.Error() == "store not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver un mensaje de confirmación
	ctx.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

// UpdateStore handles PUT requests to update an existing store

// UpdateStore godoc
// @Summary Update store
// @Description Update store
// @Tags stores
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Param store body dtos.StoreRequest true "Store to update"
// @Router /leal-test/stores/{id} [put]
func (c *StoreController) UpdateStore(ctx *gin.Context) {
	// Obtener el ID desde la URL
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
		return
	}

	// Vincular el cuerpo de la solicitud al DTO
	var storeDTO dtos.StoreRequest
	if err := ctx.ShouldBindJSON(&storeDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Mapear el DTO a los datos del modelo
	storeData := adapters.ToStoreModel(storeDTO)

	// Llamar al servicio para actualizar la tienda
	err = c.service.UpdateStore(uint(id), &storeData)
	if err != nil {
		if err.Error() == "store not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enviar una respuesta exitosa
	ctx.JSON(http.StatusOK, gin.H{"message": "Store updated successfully"})
}

// CreateStore handles POST requests to create a new store

// CreateStore godoc
// @Summary Create store
// @Description Create store
// @Tags stores
// @Accept  json
// @Produce  json
// @Param store body dtos.StoreRequest true "Store to create"
// @Router /leal-test/stores [post]
func (c *StoreController) CreateStore(ctx *gin.Context) {
	var store models.Store
	if err := ctx.ShouldBindJSON(&store); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := c.service.CreateStore(&store)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Store created successfully"})
}
