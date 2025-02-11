package models

type DebitCard struct {
	CardID    string  `gorm:"primaryKey;column:card_id"`
	UserID    *string `gorm:"column:user_id"`
	Name      *string `gorm:"column:name"`
	DummyCol7 *string `gorm:"column:dummy_col_7"`
}
