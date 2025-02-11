package models

type DebitCardDesign struct {
	CardID      string  `gorm:"primaryKey" json:"card_id"`
	UserID      *string `json:"user_id"`
	Color       *string `json:"color"`
	BorderColor *string `json:"border_color"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}
