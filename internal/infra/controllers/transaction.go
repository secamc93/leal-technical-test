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

// TransactionController struct
type TransactionController struct {
	service services.TransactionService
}

// NewTransactionController constructor
func NewTransactionController() *TransactionController {
	db := config.NewPostgresConnection()
	repo := repository.NewTransactionRepository(db)
	repobranch := repository.NewBranchRepository(db)
	repoCampaign := repository.NewCampaignRepository(db)
	service := services.NewTransactionService(repo, repobranch, repoCampaign)

	return &TransactionController{
		service: service,
	}
}

// GetAllTransactions handles GET requests to retrieve all transactions
// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /leal-test/transactions [get]
func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.service.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionsDTOs := adapters.ToTransactionDTOs(transactions)

	ctx.JSON(http.StatusOK, transactionsDTOs)
}

// GetTransactionById handles GET requests to retrieve a transaction by its ID
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Transaction ID"
// @Router /leal-test/transactions/{id} [get]
func (c *TransactionController) GetTransactionById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := c.service.GetTransactionById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionDTOs := adapters.ToTransactionDTO(transaction)

	ctx.JSON(http.StatusOK, transactionDTOs)
}

// GetTransactionsByUserId handles GET requests to retrieve transactions by UserID
// @Summary Get transactions by UserID
// @Description Get transactions by UserID
// @Tags transactions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_id path int true "User ID"
// @Router /leal-test/transactions/user/{user_id} [get]
func (c *TransactionController) GetTransactionsByUserId(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	transactions, err := c.service.GetTransactionsByUserId(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionsDTOs := adapters.ToTransactionDTOs(transactions)

	ctx.JSON(http.StatusOK, transactionsDTOs)
}

// CreateTransaction handles POST requests to create a new transaction
// @Summary Create a new transaction
// @Description Create a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param transaction body dtos.TransactionRequest true "Transaction data"
// @Router /leal-test/transactions [post]
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transactionDTO = dtos.TransactionRequest{}

	if err := ctx.ShouldBindJSON(&transactionDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := adapters.ToTransactionModel(transactionDTO)
	point, err := c.service.CreateTransaction(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"point": point})
}
