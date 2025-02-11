package models

type Banner struct {
	BannerID    string  `gorm:"primaryKey" json:"bannerId"`
	UserID      *string `json:"userId,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Image       *string `json:"image,omitempty"`
}
