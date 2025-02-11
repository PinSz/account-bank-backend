package repositories

import (
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// ดึงรายการธุรกรรมตาม userID
func (r *TransactionRepository) GetTransactionByUserID(userId string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Select("name, image, is_bank").
		Where("user_id = ?", userId).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
