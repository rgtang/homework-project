package middlewares

import (
	"my-go-project/src/pkg"
	"my-go-project/src/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GlobalRecovery 捕获程序运行时致命的 Panic
func GlobalRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录崩溃栈帧日志
				pkg.Logger.Error("[System Panic Recovered]",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
				)

				// 响应统一的 500 JSON 给前端，隐藏内部堆栈细节
				utils.RenderError(c, utils.ErrInternalServer)
				c.Abort()
			}
		}()
		c.Next()
	}
}
