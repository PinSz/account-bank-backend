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

// 🔹 เพิ่มฟังก์ชัน GetAccountInfo
func (ctrl *AccountController) GetAccountDetails(c *gin.Context) {
	var request struct {
		UserID string `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	accounts, err := ctrl.accountService.GetAccountDetails(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
