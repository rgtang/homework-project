package controllers

import (
	"net/http"

	"my-go-project/src/models"
	"my-go-project/src/pkg"
	"my-go-project/src/utils"

	"github.com/gin-gonic/gin"
)

type PostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}

	if err := pkg.DB.Create(&post).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to create post")
		return
	}

	utils.Success(c, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	// Preload 加载关联的用户信息（隐藏密码字段）
	if err := pkg.DB.Preload("User").Find(&posts).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to fetch posts")
		return
	}

	utils.Success(c, posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pkg.DB.Preload("User").First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "post not found")
		return
	}

	utils.Success(c, post)
}

func UpdatePost(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var post models.Post
	if err := pkg.DB.First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "post not found")
		return
	}

	// 鉴权：只有作者本人可以修改
	if post.UserID != userID {
		utils.Error(c, http.StatusForbidden, "unauthorized to edit this post")
		return
	}

	var input PostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	pkg.DB.Model(&post).Updates(models.Post{Title: input.Title, Content: input.Content})
	utils.Success(c, post)
}

func DeletePost(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var post models.Post
	if err := pkg.DB.First(&post, id).Error; err != nil {
		utils.Error(c, http.StatusNotFound, "post not found")
		return
	}

	// 鉴权：只有作者本人可以删除
	if post.UserID != userID {
		utils.Error(c, http.StatusForbidden, "unauthorized to delete this post")
		return
	}

	pkg.DB.Delete(&post)
	utils.Success(c, "post deleted successfully")
}
