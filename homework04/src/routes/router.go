package routes

import (
	"my-go-project/src/controllers"
	"my-go-project/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 开放路由
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.GET("/posts/:id/comments", controllers.GetCommentsByPost)
	}

	// 需认证路由
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
