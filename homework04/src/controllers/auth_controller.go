package controllers

import (
	"net/http"

	"my-go-project/src/models"
	"my-go-project/src/pkg"
	"my-go-project/src/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to hash password")
		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
		Email:    input.Email,
	}

	if err := pkg.DB.Create(&user).Error; err != nil {
		utils.Error(c, http.StatusBadRequest, "username or email already exists")
		return
	}

	utils.Success(c, "registration successful")
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := pkg.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.Error(c, http.StatusUnauthorized, "invalid username or password")
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.Error(c, http.StatusUnauthorized, "invalid username or password")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to generate token")
		return
	}

	utils.Success(c, gin.H{"token": token})
}
