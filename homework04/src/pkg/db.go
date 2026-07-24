package pkg

import (
	"my-go-project/src/models"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		Logger.Fatal("Failed to connect database", zap.Error(err))
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		Logger.Fatal("Auto migrate failed", zap.Error(err))
	}

	Logger.Info("Database initialized successfully")
}
