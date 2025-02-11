package services

import (
	"account-bank-backend/repositories"
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
	BorderColor *string `json:"border_color"`
}

func (s *DebitCardService) GetDebitCardInfo(userID string) (*DebitCardResponse, error) {
	card, _ := s.Repo.GetDebitCardByUserID(userID)
	status, _ := s.Repo.GetDebitCardStatusByUserID(userID)
	detail, _ := s.Repo.GetDebitCardDetailByUserID(userID)
	design, _ := s.Repo.GetDebitCardDesignByUserID(userID)

	return &DebitCardResponse{
		Name:        card.Name,
		Status:      status.Status,
		Issuer:      detail.Issuer,
		Number:      detail.Number,
		Color:       design.Color,
		BorderColor: design.BorderColor,
	}, nil
}
