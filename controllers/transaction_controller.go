package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: transactionService}
}

func (tc *TransactionController) GetTransactionsByUserID(c *gin.Context) {
	userId := c.Param("userId") // ดึง userId จาก URL

	transactions, err := tc.TransactionService.GetTransactionByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
