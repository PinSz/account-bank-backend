package services

import (
	"account-bank-backend/repositories"
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
	return s.Repo.GetUserNameByID(userId)
}
