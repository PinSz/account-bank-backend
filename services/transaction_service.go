package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
	"fmt"
)

type TransactionService struct {
	Repo *repositories.TransactionRepository
}

// สร้าง instance ของ TransactionService โดยรับ TransactionRepository
func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{Repo: repo}
}

// ดึงรายการธุรกรรมตาม userID
func (s *TransactionService) GetTransactionByUserID(userId string) ([]models.Transaction, error) {
	transactions, err := s.Repo.GetTransactionByUserID(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions for user")
	}
	return transactions, nil
}
