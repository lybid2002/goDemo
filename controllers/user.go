package controllers

import (
	"errors"
	"gotest/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.UserInput
	// 绑定并进行验证
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := models.GetErrorMsgs(ve)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
	})
}

var userController = new(UserController)

func GetUserController() *UserController {
	return userController
}
