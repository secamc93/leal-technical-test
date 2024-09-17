package router

import (
	"leal-technical-test/config"
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
		// Public routes
		lealTestGroup.POST("/login", r.userController.Login)
		lealTestGroup.POST("/users", r.userController.CreateUser)

		// Protected routes
		protected := lealTestGroup.Group("/")
		protected.Use(config.NewTokenManager().AuthMiddleware())
		{
			// Store routes
			protected.GET("/stores", r.storeController.GetAllStores)
			protected.GET("/stores/:id", r.storeController.GetStoreById)
			protected.DELETE("/stores/:id", r.storeController.DeleteStore)
			protected.PUT("/stores/:id", r.storeController.UpdateStore)
			protected.POST("/stores", r.storeController.CreateStore)

			protected.GET("/users", r.userController.GetAllUsers)
			protected.GET("/users/:id", r.userController.GetUserById)
			protected.DELETE("/users/:id", r.userController.DeleteUser)
			protected.PUT("/users/:id", r.userController.UpdateUser)

			protected.GET("/branches", r.branchController.GetAllBranches)
			protected.GET("/branches/:id", r.branchController.GetBranchById)
			protected.DELETE("/branches/:id", r.branchController.DeleteBranch)
			protected.PUT("/branches/:id", r.branchController.UpdateBranch)
			protected.POST("/branches", r.branchController.CreateBranch)

			protected.GET("/campaigns", r.campaignController.GetAllCampaigns)
			protected.GET("/campaigns/:id", r.campaignController.GetCampaignById)
			protected.POST("/campaigns", r.campaignController.CreateCampaign)
			protected.PUT("/campaigns/:id", r.campaignController.UpdateCampaign)
			protected.DELETE("/campaigns/:id", r.campaignController.DeleteCampaign)

			protected.GET("/acumulaterewards", r.accumulatedRewardController.GetAllRewards)
			protected.GET("/acumulaterewards/:id", r.accumulatedRewardController.GetRewardById)
			protected.GET("/acumulaterewards/user/:user_id/store/:store_id", r.accumulatedRewardController.GetRewardByUserAndStore)

			protected.GET("/rewards", r.rewardController.GetAllRewards)
			protected.GET("/rewards/:id", r.rewardController.GetRewardById)
			protected.GET("/rewards/store/:store_id", r.rewardController.GetRewardsByStoreId)
			protected.POST("/rewards", r.rewardController.CreateReward)
			protected.PUT("/rewards/:id", r.rewardController.UpdateReward)
			protected.DELETE("/rewards/:id", r.rewardController.DeleteReward)
			protected.GET("/rewards/claim/:user_id/:reward_id/:store_id", r.rewardController.GetClaimRewardPoints)

			protected.GET("/transactions", r.transactionController.GetAllTransactions)
			protected.GET("/transactions/:id", r.transactionController.GetTransactionById)
			protected.GET("/transactions/user/:user_id", r.transactionController.GetTransactionsByUserId)
			protected.POST("/transactions", r.transactionController.CreateTransaction)
		}
	}
}
