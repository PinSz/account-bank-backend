package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannerController struct {
	BannerService *services.BannerService
}

func NewBannerController(bannerService *services.BannerService) *BannerController {
	return &BannerController{BannerService: bannerService}
}

// ✅ API GET /api/banners/:userId
func (c *BannerController) GetBannersByUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	// เรียกใช้ Service
	banners, err := c.BannerService.GetBannersByUserID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get banners"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"banners": banners})
}
