package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
)

type BannerService struct {
	Repo *repositories.BannerRepository
}

// สร้าง instance ของ BannerService โดยรับ BannerRepository
func NewBannerService(repo *repositories.BannerRepository) *BannerService {
	return &BannerService{Repo: repo}
}

// ✅ ดึง Banner ตาม userId
func (s *BannerService) GetBannerByUserID(userId string) (*models.Banner, error) {
	return s.Repo.GetBannerByUserID(userId)
}
