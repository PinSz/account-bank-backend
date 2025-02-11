package models

type AccountDetail struct {
	AccountID     string  `gorm:"primaryKey" json:"account_id"`
	UserID        *string `json:"user_id"`
	Color         *string `json:"color"`
	IsMainAccount *bool   `json:"is_main_account"`
	Progress      *int    `json:"progress"`
}
