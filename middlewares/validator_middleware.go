package middlewares

import (
	"github.com/gin-gonic/gin"
)

/*
中间件：记录请求处理时间
*/
func ValidatorMiddleware(c *gin.Context) {
	// 在上下文中加入验证器
	// c.Set("validate", models.Validate)

	// 处理请求
	c.Next()
}
