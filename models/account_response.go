package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Struct สำหรับ JSONB
type Flag struct {
	FlagType  string `json:"flagType"`
	FlagValue string `json:"flagValue"`
	CreatedAt string `json:"createdAt"`
}

type Flags []Flag

// Scan JSONB จาก DB
func (j *Flags) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan *Flags")
	}
	return json.Unmarshal(bytes, j)
}

// Value แปลง JSONB เป็น JSON
func (j Flags) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Model สำหรับ Response
type AccountResponse struct {
	AccountID     string  `gorm:"column:account_id" json:"accountId,omitempty"`
	AccountNumber string  `gorm:"column:account_number" json:"accountNumber,omitempty"`
	Color         string  `gorm:"column:color" json:"color,omitempty"`
	IsMainAccount bool    `gorm:"column:is_main_account" json:"isMainAccount,omitempty"`
	Progress      int     `gorm:"column:progress" json:"progress,omitempty"`
	Amount        float64 `gorm:"column:amount" json:"amount,omitempty"`
	Flags         Flags   `gorm:"column:flags" json:"flags,omitempty"`
}
