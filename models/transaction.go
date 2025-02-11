package models

type Transaction struct {
	TransactionID string  `gorm:"primaryKey;column:transaction_id" json:"transactionId,omitempty"`
	UserID        *string `gorm:"column:user_id" json:"userId,omitempty"`
	Name          *string `gorm:"column:name" json:"name,omitempty"`
	Image         *string `gorm:"column:image" json:"image,omitempty"`
	IsBank        *bool   `gorm:"column:is_bank" json:"isBank,omitempty"`
	DummyCol6     *string `gorm:"column:dummy_col_6" json:"dummyCol6,omitempty"`
}
