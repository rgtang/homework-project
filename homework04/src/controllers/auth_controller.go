package controllers

import (
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
		utils.RenderError(c, utils.ErrInvalidParam.WithErr(err))
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.RenderError(c, utils.ErrInternalServer.WithErr(err))
		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
		Email:    input.Email,
	}

	if err := pkg.DB.Create(&user).Error; err != nil {
		utils.RenderError(c, utils.ErrUserExists.WithErr(err))
		return
	}

	utils.Success(c, "registration successful")
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RenderError(c, utils.ErrInvalidParam.WithErr(err))
		return
	}

	var user models.User
	if err := pkg.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.RenderError(c, utils.ErrLoginFailed)
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.RenderError(c, utils.ErrLoginFailed)
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.RenderError(c, utils.ErrInternalServer.WithErr(err))
		return
	}

	utils.Success(c, gin.H{"token": token})
}
