package repositories

import (
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type BannerRepository struct {
	DB *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
	return &BannerRepository{DB: db}
}

// ✅ ดึง Banner ตาม userId
func (r *BannerRepository) GetBannerByUserID(userId string) (*models.Banner, error) {
	var banner models.Banner
	err := r.DB.Where("user_id = ?", userId).Take(&banner).Error
	if err != nil {
		return nil, err
	}
	return &banner, nil
}
