package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
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
	return s.Repo.GetTransactionByUserID(userId)
}
