package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (ctrl *AccountController) GetAccountDetails(c *gin.Context) {
	var request struct {
		UserID string `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid request",
		})
		return
	}

	accounts, err := ctrl.accountService.GetAccountDetails(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Failed to fetch data",
		})
		return
	}

	// ถ้า accounts ว่าง ให้ return 400
	if accounts == nil || len(accounts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "No account data found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   accounts,
	})
}
