package routes

import (
	"account-bank-backend/controllers"
	"account-bank-backend/repositories"
	"account-bank-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {
	// 🔹 สร้าง Repository
	accountRepo := repositories.NewAccountRepository(db)
	debitCardRepo := repositories.NewDebitCardRepository(db)
	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	bannerRepo := repositories.NewBannerRepository(db)

	// 🔹 สร้าง Service
	accountService := services.NewAccountService(accountRepo)
	debitCardService := services.NewDebitCardService(debitCardRepo)
	userService := services.NewUserService(userRepo)
	transactionService := services.NewTransactionService(transactionRepo)
	bannerService := services.NewBannerService(bannerRepo)

	// 🔹 สร้าง Controller
	accountController := controllers.NewAccountController(accountService)
	debitCardController := controllers.NewDebitCardController(debitCardService)
	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionController(transactionService)
	bannerController := controllers.NewBannerController(bannerService)

	api := router.Group("/api")
	{
		api.POST("/account/details", accountController.GetAccountDetails)
		api.GET("/user/:userId", userController.GetUserName)
		api.GET("/transactions/:userId", transactionController.GetTransactionsByUserID)
		api.GET("/banner/:userId", bannerController.GetBannersByUser)
		api.POST("/debit-card/info", debitCardController.GetDebitCardInfo)
	}
}
