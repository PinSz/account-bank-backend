package controllers

import (
	"account-bank-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

// สร้าง instance ของ UserController
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// ดึงชื่อของ user ตาม userID
func (uc *UserController) GetUserName(c *gin.Context) {
	userId := c.Param("userId") // ดึง userID จาก URL parameter

	name, err := uc.UserService.GetUserNameByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"userId": userId,
			"name":   name,
		},
	})
}
