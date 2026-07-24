package controllers

import (
	"strconv"

	"my-go-project/src/models"
	"my-go-project/src/pkg"
	"my-go-project/src/utils"

	"github.com/gin-gonic/gin"
)

type CommentInput struct {
	Content string `json:"content" binding:"required"`
}

func CreateComment(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		utils.RenderError(c, utils.ErrInvalidParam.WithErr(err))
		return
	}

	// 检查文章是否存在
	var post models.Post
	if err := pkg.DB.First(&post, postID).Error; err != nil {
		utils.RenderError(c, utils.ErrNotFoundPost)
		return
	}

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RenderError(c, utils.ErrInvalidParam.WithErr(err))
		return
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  uint(postID),
	}

	if err := pkg.DB.Create(&comment).Error; err != nil {
		utils.RenderError(c, utils.ErrDatabaseError.WithErr(err))
		return
	}

	utils.Success(c, comment)
}

func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")

	var comments []models.Comment
	if err := pkg.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.RenderError(c, utils.ErrDatabaseError.WithErr(err))
		return
	}

	utils.Success(c, comments)
}
