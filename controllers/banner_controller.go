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
	banner, err := c.BannerService.GetBannerByUserID(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Failed to get banner",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"banner": banner,
	})
}
