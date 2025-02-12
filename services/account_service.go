package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
	"fmt"
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
	accounts, err := s.accountRepo.GetAccountDetails(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get account details for user %s: %w", userID, err)
	}
	return accounts, nil
}
