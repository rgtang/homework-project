package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, ResponseData{
		Code:    httpCode,
		Message: message,
	})
}
