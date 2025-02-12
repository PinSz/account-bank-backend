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

	// ✅ ตรวจสอบว่าข้อมูลที่รับมาถูกต้องหรือไม่
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid request",
		})
		return
	}

	// ✅ ดึงข้อมูลจาก Service
	data, err := c.Service.GetDebitCardInfo(request.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Failed to retrieve debit card info",
		})
		return
	}

	// ✅ ส่งข้อมูลกลับตามรูปแบบที่ต้องการ
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}
