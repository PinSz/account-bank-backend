package repositories

import (
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type DebitCardRepository struct {
	DB *gorm.DB
}

func NewDebitCardRepository(db *gorm.DB) *DebitCardRepository {
	return &DebitCardRepository{DB: db}
}

func (r *DebitCardRepository) GetDebitCardByUserID(userID string) (*models.DebitCard, error) {
	var card models.DebitCard
	err := r.DB.Select("name").
		Where("user_id = ?", userID).
		Take(&card).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (r *DebitCardRepository) GetDebitCardStatusByUserID(userID string) (*models.DebitCardStatus, error) {
	var status models.DebitCardStatus
	err := r.DB.Select("status").
		Where("user_id = ?", userID).
		Take(&status).Error
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *DebitCardRepository) GetDebitCardDetailByUserID(userID string) (*models.DebitCardDetail, error) {
	var detail models.DebitCardDetail
	err := r.DB.Select("issuer, number").
		Where("user_id = ?", userID).
		Take(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

func (r *DebitCardRepository) GetDebitCardDesignByUserID(userID string) (*models.DebitCardDesign, error) {
	var design models.DebitCardDesign
	err := r.DB.Select("color, border_color").
		Where("user_id = ?", userID).
		Take(&design).Error
	if err != nil {
		return nil, err
	}
	return &design, nil
}
