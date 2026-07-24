package routes

import (
	"my-go-project/src/controllers"
	"my-go-project/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 使用 gin.New() 而非 gin.Default()，防止重复挂载 Gin 默认的 Logger/Recovery
	r := gin.New()

	// 挂载我们自定义的 Zap 日志与 Panic 恢复中间件
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.GlobalRecovery())

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.GET("/posts/:id/comments", controllers.GetCommentsByPost)
	}

	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuth())
	{
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)

		protected.POST("/posts/:id/comments", controllers.CreateComment)
	}

	return r
}
