package utils

import (
	"my-go-project/src/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, ResponseData{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error 快捷响应：直接传入 HTTP 状态码和消息字符串
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ResponseData{
		Code:    statusCode,
		Message: message,
	})
}

// RenderError 统一渲染 AppError 并记录日志
func RenderError(c *gin.Context, appErr *AppError) {
	// 如果包含底层的系统/数据库错误，记录 ERROR 级别日志方便排查
	if appErr.RawErr != nil {
		pkg.Logger.Error("[Business Error]",
			zap.String("path", c.Request.URL.Path),
			zap.Int("code", appErr.Code),
			zap.Error(appErr.RawErr),
		)
	}

	c.JSON(appErr.StatusCode, ResponseData{
		Code:    appErr.Code,
		Message: appErr.Message,
	})
}
