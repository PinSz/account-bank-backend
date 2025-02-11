package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
)

type AccountService interface {
	GetAccountDetails(userID string) ([]models.AccountResponse, error)
}

type accountService struct {
	accountRepo repositories.AccountRepository
}

func NewAccountService(accountRepo repositories.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) GetAccountDetails(userID string) ([]models.AccountResponse, error) {
	return s.accountRepo.GetAccountDetails(userID)
}
