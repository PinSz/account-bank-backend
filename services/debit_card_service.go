package services

import (
	"account-bank-backend/repositories"
	"fmt"
)

type DebitCardService struct {
	Repo *repositories.DebitCardRepository
}

func NewDebitCardService(repo *repositories.DebitCardRepository) *DebitCardService {
	return &DebitCardService{Repo: repo}
}

type DebitCardResponse struct {
	Name        *string `json:"name"`
	Status      *string `json:"status"`
	Issuer      *string `json:"issuer"`
	Number      *string `json:"number"`
	Color       *string `json:"color"`
	BorderColor *string `json:"borderColor"`
}

func (s *DebitCardService) GetDebitCardInfo(userID string) (*DebitCardResponse, error) {
	card, err := s.Repo.GetDebitCardByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get debit card: %w", err)
	}

	status, err := s.Repo.GetDebitCardStatusByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get debit card status: %w", err)
	}

	detail, err := s.Repo.GetDebitCardDetailByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get debit card detail: %w", err)
	}

	design, err := s.Repo.GetDebitCardDesignByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get debit card design: %w", err)
	}

	return &DebitCardResponse{
		Name:        card.Name,
		Status:      status.Status,
		Issuer:      detail.Issuer,
		Number:      detail.Number,
		Color:       design.Color,
		BorderColor: design.BorderColor,
	}, nil
}
