package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountBalanceController struct {
	Service *services.AccountBalanceService
}

func NewAccountBalanceController(service *services.AccountBalanceService) *AccountBalanceController {
	return &AccountBalanceController{Service: service}
}

func (c *AccountBalanceController) GetAllBalances(ctx *gin.Context) {
	balances, err := c.Service.GetAllBalances()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, balances)
}

// func (ctrl *AccountBalanceController) GetAllBalances(c *gin.Context) {
// 	// ส่งค่า 200 OK กลับมา พร้อมกับ JSON response
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Success",
// 		"data":    []interface{}{},
// 	})
// }
