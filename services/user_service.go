package services

import (
	"account-bank-backend/repositories"
	"fmt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// สร้าง instance ของ UserService โดยรับ UserRepository
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// ดึงชื่อของ user ตาม userID
func (s *UserService) GetUserNameByID(userId string) (string, error) {
	userName, err := s.Repo.GetUserNameByID(userId)
	if err != nil {
		return "", fmt.Errorf("failed to get user name for user")
	}
	return userName, nil
}
