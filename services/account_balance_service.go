package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
)

type AccountBalanceService struct {
	Repo *repositories.AccountBalanceRepository
}

func NewAccountBalanceService(repo *repositories.AccountBalanceRepository) *AccountBalanceService {
	return &AccountBalanceService{Repo: repo}
}

func (s *AccountBalanceService) GetAllBalances() ([]models.AccountBalance, error) {
	return s.Repo.GetAll()
}
