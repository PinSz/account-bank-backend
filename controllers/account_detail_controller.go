package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountDetailController struct {
	Service *services.AccountDetailService
}

func NewAccountDetailController(service *services.AccountDetailService) *AccountDetailController {
	return &AccountDetailController{Service: service}
}

func (c *AccountDetailController) GetAllDetails(ctx *gin.Context) {
	details, err := c.Service.GetAllDetails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, details)
}
