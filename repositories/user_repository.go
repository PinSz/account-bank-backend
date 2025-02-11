package repositories

import (
	"account-bank-backend/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ดึงชื่อของ user ตาม userID
func (r *UserRepository) GetUserNameByID(userId string) (string, error) {
	var name string
	err := r.DB.Model(&models.User{}).Select("name").Where("user_id = ?", userId).Scan(&name).Error
	if err != nil {
		return "", err
	}
	if name == "" {
		return "", errors.New("user not found or name is NULL")
	}
	return name, nil
}
