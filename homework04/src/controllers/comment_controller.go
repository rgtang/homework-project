package controllers

import (
	"net/http"
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
		utils.Error(c, http.StatusBadRequest, "invalid post id")
		return
	}

	// 检查文章是否存在
	var post models.Post
	if err := pkg.DB.First(&post, postID).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "post not found")
		return
	}

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  uint(postID),
	}

	if err := pkg.DB.Create(&comment).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to create comment")
		return
	}

	utils.Success(c, comment)
}

func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")

	var comments []models.Comment
	if err := pkg.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to fetch comments")
		return
	}

	utils.Success(c, comments)
}
