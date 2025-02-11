package models

type DebitCardDetail struct {
	CardID string  `gorm:"primaryKey" json:"card_id"`
	UserID *string `json:"user_id"`
	Issuer *string `json:"issuer"`
	Number *string `json:"number"`
}
