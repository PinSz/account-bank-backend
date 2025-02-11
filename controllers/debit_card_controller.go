package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DebitCardController struct {
	Service *services.DebitCardService
}

func NewDebitCardController(service *services.DebitCardService) *DebitCardController {
	return &DebitCardController{Service: service}
}

func (c *DebitCardController) GetDebitCardInfo(ctx *gin.Context) {
	var request struct {
		UserID string `json:"userId" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	data, err := c.Service.GetDebitCardInfo(request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
