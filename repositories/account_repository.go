package repositories

import (
	"account-bank-backend/config"
	"account-bank-backend/models"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountDetails(userID string) ([]models.AccountResponse, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) GetAccountDetails(userID string) ([]models.AccountResponse, error) {
	println("userID--->", userID)
	db := config.DB
	var accounts []models.AccountResponse

	err := db.Raw(`
		SELECT 
			a.account_id, 
			a.account_number, 
			ad.color, 
			ad.is_main_account, 
			ad.progress, 
			ab.amount,
			COALESCE(
				jsonb_agg(
					jsonb_build_object(
						'flagType', af.flag_type,
						'flagValue', af.flag_value,
						'createdAt', af.created_at
					)
				) FILTER (WHERE af.flag_type IS NOT NULL),
				'[]'::jsonb
			) AS flags
		FROM accounts a
		LEFT JOIN account_details ad ON a.account_id = ad.account_id
		LEFT JOIN account_balances ab ON a.account_id = ab.account_id
		LEFT JOIN account_flags af ON a.account_id = af.account_id
		WHERE a.user_id = ?
		GROUP BY a.account_id, a.account_number, ad.color, ad.is_main_account, ad.progress, ab.amount;
	`, userID).Scan(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
