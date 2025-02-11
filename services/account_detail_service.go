package services

import (
	"account-bank-backend/models"
	"account-bank-backend/repositories"
)

type AccountDetailService struct {
	Repo *repositories.AccountDetailRepository
}

func NewAccountDetailService(repo *repositories.AccountDetailRepository) *AccountDetailService {
	return &AccountDetailService{Repo: repo}
}

func (s *AccountDetailService) GetAllDetails() ([]models.AccountDetail, error) {
	return s.Repo.GetAll()
}
