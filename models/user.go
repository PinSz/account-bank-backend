package models

type User struct {
	UserID    string  `gorm:"primaryKey;column:user_id" json:"userId,omitempty"`
	Name      *string `gorm:"column:name" json:"name,omitempty"`
	DummyCol1 *string `gorm:"column:dummy_col_1" json:"dummyCol1,omitempty"`
}
