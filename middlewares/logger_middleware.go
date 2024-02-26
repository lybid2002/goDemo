package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 获取返回的body内容
type ExResponseWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (writer ExResponseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	writer.bodyBuf.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return writer.ResponseWriter.Write(b)
}

// JSONToPlainString 将JSON格式的字符串转换为普通字符串
func Json2PlainString(jsonStr string) (string, error) {
	var raw json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &raw)

	if err != nil {
		return "", fmt.Errorf("解析JSON字符串时出错: %w", err)
	}

	return string(raw), nil
}

func LoggerTestMiddleware(c *gin.Context) {
	// 获取请求方法
	method := c.Request.Method
	// url := c.Request.URL.Path

	// 创建对象，用于记录请求体
	writer := ExResponseWriter{c.Writer, bytes.NewBuffer([]byte{})}
	// 设置Writer为自定义的ResponseWriter
	c.Writer = writer

	// 继续处理请求
	c.Next()

	// 如果请求方法不是GET
	if method != "GET" {
		// 在goroutine中打印请求体
		go func(source string) {
			fmt.Printf("Response body %s\n", source)
		}(writer.bodyBuf.String())
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始
		// logrus.Infof("Request: %s %s", c.Request.Method, c.Request.URL)
		requestInfo := fmt.Sprintf("Request: %s %s", c.Request.Method, c.Request.URL)

		// 创建exResponseWriter对象，用于记录请求体
		writer := ExResponseWriter{c.Writer, bytes.NewBuffer([]byte{})}
		// 设置Writer为自定义的ResponseWriter
		c.Writer = writer

		// 处理请求
		c.Next()

		// 记录请求结束
		strBody := writer.bodyBuf.String()

		// logrus.Infof("Response: %d", c.Writer.Status())
		responseInfo := fmt.Sprintf("Response: %d body: %s", c.Writer.Status(), strBody)
		// fmt.Println(responseInfo)
		logrus.Infof("%s \n%s", requestInfo, responseInfo)
	}
}
