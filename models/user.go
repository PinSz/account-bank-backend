package models

type User struct {
	UserID    string  `gorm:"primaryKey;column:user_id"`
	Name      *string `gorm:"column:name"`
	DummyCol1 *string `gorm:"column:dummy_col_1"`
}
