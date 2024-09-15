package router

import (
	"leal-technical-test/internal/infra/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router struct
type Router struct {
	engine                      *gin.Engine
	storeController             *controllers.StoreController
	userController              *controllers.UserController
	branchController            *controllers.BranchController
	campaignController          *controllers.CampaignController
	accumulatedRewardController *controllers.AccumulatedRewardController
	rewardController            *controllers.RewardController
	transactionController       *controllers.TransactionController
}

// NewRouter constructor
func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine:                      engine,
		storeController:             controllers.NewStoreController(),
		userController:              controllers.NewUserController(),
		branchController:            controllers.NewBranchController(),
		campaignController:          controllers.NewCampaignController(),
		accumulatedRewardController: controllers.NewAccumulatedRewardController(),
		rewardController:            controllers.NewRewardController(),
		transactionController:       controllers.NewTransactionController(),
	}
}

// InitializeRoutes sets up the routes for the application
func (r *Router) InitializeRoutes() {
	// Swagger route
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	lealTestGroup := r.engine.Group("/leal-test")
	{
		// Store routes
		lealTestGroup.GET("/stores", r.storeController.GetAllStores)
		lealTestGroup.GET("/stores/:id", r.storeController.GetStoreById)
		lealTestGroup.DELETE("/stores/:id", r.storeController.DeleteStore)
		lealTestGroup.PUT("/stores/:id", r.storeController.UpdateStore)
		lealTestGroup.POST("/stores", r.storeController.CreateStore)

		lealTestGroup.GET("/users", r.userController.GetAllUsers)
		lealTestGroup.GET("/users/:id", r.userController.GetUserById)
		lealTestGroup.DELETE("/users/:id", r.userController.DeleteUser)
		lealTestGroup.PUT("/users/:id", r.userController.UpdateUser)
		lealTestGroup.POST("/users", r.userController.CreateUser)

		lealTestGroup.GET("/branches", r.branchController.GetAllBranches)
		lealTestGroup.GET("/branches/:id", r.branchController.GetBranchById)
		lealTestGroup.DELETE("/branches/:id", r.branchController.DeleteBranch)
		lealTestGroup.PUT("/branches/:id", r.branchController.UpdateBranch)
		lealTestGroup.POST("/branches", r.branchController.CreateBranch)

		lealTestGroup.GET("/campaigns", r.campaignController.GetAllCampaigns)
		lealTestGroup.GET("/campaigns/:id", r.campaignController.GetCampaignById)
		lealTestGroup.POST("/campaigns", r.campaignController.CreateCampaign)
		lealTestGroup.PUT("/campaigns/:id", r.campaignController.UpdateCampaign)
		lealTestGroup.DELETE("/campaigns/:id", r.campaignController.DeleteCampaign)

		lealTestGroup.GET("/acumulaterewards", r.accumulatedRewardController.GetAllRewards)
		lealTestGroup.GET("/acumulaterewards/:id", r.accumulatedRewardController.GetRewardById)
		lealTestGroup.GET("/acumulaterewards/user/:user_id/store/:store_id", r.accumulatedRewardController.GetRewardByUserAndStore)
		lealTestGroup.POST("/acumulaterewards", r.accumulatedRewardController.CreateReward)
		lealTestGroup.PUT("/acumulaterewards/:id", r.accumulatedRewardController.UpdateReward)
		lealTestGroup.DELETE("/acumulaterewards/:id", r.accumulatedRewardController.DeleteReward)

		lealTestGroup.GET("/rewards", r.rewardController.GetAllRewards)
		lealTestGroup.GET("/rewards/:id", r.rewardController.GetRewardById)
		lealTestGroup.GET("/rewards/store/:store_id", r.rewardController.GetRewardsByStoreId)
		lealTestGroup.POST("/rewards", r.rewardController.CreateReward)
		lealTestGroup.PUT("/rewards/:id", r.rewardController.UpdateReward)
		lealTestGroup.DELETE("/rewards/:id", r.rewardController.DeleteReward)

		lealTestGroup.GET("/transactions", r.transactionController.GetAllTransactions)
		lealTestGroup.GET("/transactions/:id", r.transactionController.GetTransactionById)
		lealTestGroup.GET("/transactions/user/:user_id", r.transactionController.GetTransactionsByUserId)
		lealTestGroup.POST("/transactions", r.transactionController.CreateTransaction)

	}
}
