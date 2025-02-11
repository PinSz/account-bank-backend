package repositories

import (
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type AccountDetailRepository struct {
	DB *gorm.DB
}

func NewAccountDetailRepository(db *gorm.DB) *AccountDetailRepository {
	return &AccountDetailRepository{DB: db}
}

func (r *AccountDetailRepository) GetAll() ([]models.AccountDetail, error) {
	var details []models.AccountDetail
	err := r.DB.Find(&details).Error
	return details, err
}
