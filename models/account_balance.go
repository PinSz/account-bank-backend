package models

type AccountBalance struct {
	AccountID string   `gorm:"primaryKey" json:"account_id"`
	UserID    *string  `json:"user_id"`
	Amount    *float64 `json:"amount"`
}
