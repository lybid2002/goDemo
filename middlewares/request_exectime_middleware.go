package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
中间件：记录请求处理时间
*/
func RequestExectimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// c.Writer
		// 请求处理完成后
		duration := time.Since(startTime)
		fmt.Printf("Request took %v\n", duration)
	}
}
