package models

type DebitCardStatus struct {
	CardID    string  `gorm:"primaryKey;column:card_id"`
	UserID    *string `gorm:"column:user_id"`
	Status    *string `gorm:"column:status"`
	DummyCol8 *string `gorm:"column:dummy_col_8"`
}

func (DebitCardStatus) TableName() string {
	return "debit_card_status"
}
