package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			logrus.Errorf("Error: %v", err)
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": c.Errors})
		c.Abort()
	}
}
