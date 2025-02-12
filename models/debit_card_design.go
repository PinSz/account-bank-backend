package models

type DebitCardDesign struct {
	CardID      string  `gorm:"primaryKey" json:"cardId"`
	UserID      *string `json:"userId,omitempty"`
	Color       *string `json:"color,omitempty"`
	BorderColor *string `json:"borderColor,omitempty"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}
