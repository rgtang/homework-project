package main

import (
	"fmt"
	"log"

	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 1.定义模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex"`
	PostCount int    `gorm:"default:0"`
	// 一对多关系：一个用户可以有多篇文章 (Posts)
	Posts []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title         string `gorm:"type:varchar(200);not null"`
	Content       string `gorm:"type:text"`
	CommentStatus string `gorm:"type:varchar(20);default:'有评论'"` // 评论状态
	// 外键：属于哪个用户
	UserID uint
	// 一对多关系：一篇文章可以有多个评论 (Comments)
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	// 外键：属于哪篇文章
	PostID uint
	// 外键：评论发布者是谁
	UserID uint
}

// Post钩子函数：创建文章时自动增加用户的 post_count
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error
}

// Comment 钩子：删除评论后检查文章评论数，若为 0 更改状态为 "无评论"
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

// ==================== 2. 数据构造逻辑 (Seed Data) ====================

// SeedData 创建测试初始数据
func SeedData(db *gorm.DB) error {
	log.Println("🌱 开始造数据...")

	// 1. 创建 2 个测试用户
	users := []User{
		{Username: "张三", Email: "zhangsan@example.com"},
		{Username: "李四", Email: "lisi@example.com"},
	}
	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("创建用户失败: %w", err)
	}

	// 2. 为张三创建 2 篇文章 (会触发 Post 的 AfterCreate 钩子，自动加 PostCount)
	posts := []Post{
		{Title: "Go 语言从入门到放弃", Content: "Go 是一门极简的语言...", UserID: users[0].ID},
		{Title: "GORM 深度解析", Content: "GORM 是 Go 最流行的 ORM 框架...", UserID: users[0].ID},
		{Title: "李四的日常随笔", Content: "今天天气不错...", UserID: users[1].ID},
	}
	if err := db.Create(&posts).Error; err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}

	// 3. 为文章添加评论
	comments := []Comment{
		// 张三的第一篇文章有 3 条评论（成为评论最多的文章）
		{Content: "写得太好了！", UserID: users[1].ID, PostID: posts[0].ID},
		{Content: "学习了，非常实用。", UserID: users[0].ID, PostID: posts[0].ID},
		{Content: "赞同，期待下一篇！", UserID: users[1].ID, PostID: posts[0].ID},

		// 张三的第二篇文章有 1 条评论
		{Content: "GORM 的钩子真的很方便", UserID: users[1].ID, PostID: posts[1].ID},
	}
	if err := db.Create(&comments).Error; err != nil {
		return fmt.Errorf("创建评论失败: %w", err)
	}

	log.Println("✅ 测试数据造号完成")
	return nil
}

func main() {
	// 删除旧数据库文件（方便每次运行重新造数据）
	_ = os.Remove("test_blog.db")

	// 初始化数据库连接
	db, err := gorm.Open(sqlite.Open("test_blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 设置为 Silent 减少控制台杂音
	})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 1. 自动迁移建表
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatalf("建表失败: %v", err)
	}

	// 2. 插入测试数据
	if err := SeedData(db); err != nil {
		log.Fatalf("造数据失败: %v", err)
	}

	// ------------------ 验证 1：查看 Post 后置钩子效果 ------------------
	var zhangsan User
	db.First(&zhangsan, "username = ?", "张三")
	fmt.Printf("[钩子验证] 张三的 PostCount 自动更新为: %d (预想值: 2)\n\n", zhangsan.PostCount)

	// ------------------ 验证 2：查询某个用户的所有文章及评论 ------------------
	var userWithPosts User
	db.Preload("Posts.Comments").First(&userWithPosts, zhangsan.ID)

	fmt.Printf("--- 查询用户 [%s] 发布的所有文章及评论 ---\n", userWithPosts.Username)
	for _, post := range userWithPosts.Posts {
		fmt.Printf("文章标题: 《%s》 (评论数: %d)\n", post.Title, len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("  └─ 评论内容: %s\n", comment.Content)
		}
	}
	fmt.Println()

	// ------------------ 验证 3：查询评论最多的文章 ------------------
	var topPost Post
	db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id AND comments.deleted_at IS NULL").
		Where("posts.deleted_at IS NULL").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Preload("Comments").
		First(&topPost)

	fmt.Printf("--- 评论最多的文章 ---\n标题: 《%s》, 评论总数: %d 条\n\n", topPost.Title, len(topPost.Comments))

	// ------------------ 验证 4：删除评论测试 Comment 后置钩子 ------------------
	fmt.Println("--- 测试删除评论钩子 ---")
	// 找到张三第二篇文章的那条唯一评论并删除它
	var singleComment Comment
	db.Where("post_id = ?", userWithPosts.Posts[1].ID).First(&singleComment)

	fmt.Printf("准备删除文章 《%s》 的唯一一条评论 (ID: %d)...\n", userWithPosts.Posts[1].Title, singleComment.ID)
	db.Delete(&singleComment) // 触发 Comment 的 AfterDelete 钩子

	// 重新检查该文章的 comment_status
	var updatedPost Post
	db.First(&updatedPost, userWithPosts.Posts[1].ID)
	fmt.Printf("删除后文章的 CommentStatus 自动变更为: \"%s\"\n", updatedPost.CommentStatus)
}
