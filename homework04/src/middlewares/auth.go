package middlewares

import (
	"net/http"
	"strings"

	"my-go-project/src/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "authorization header required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.Error(c, http.StatusUnauthorized, "invalid authorization format")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		// 将当前登录用户 ID 存入上下文
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
