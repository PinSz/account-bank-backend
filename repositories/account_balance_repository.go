package repositories

import (
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type AccountBalanceRepository struct {
	DB *gorm.DB
}

func NewAccountBalanceRepository(db *gorm.DB) *AccountBalanceRepository {
	return &AccountBalanceRepository{DB: db}
}

func (r *AccountBalanceRepository) GetAll() ([]models.AccountBalance, error) {
	var balances []models.AccountBalance
	err := r.DB.Find(&balances).Error
	return balances, err
}
