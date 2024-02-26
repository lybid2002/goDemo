package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var FiledTagFunc = map[string]func(validator.FieldError) string{
	"uniqueEmail": func(fe validator.FieldError) string {
		return fmt.Sprintf("%s值%s已经被使用", fe.Field(), fe.Value())
	},
	"testTagName": func(fe validator.FieldError) string {
		logrus.Debugf("test tag name func")
		return "test tag name"
	},
}

func GetErrorMsg(fe validator.FieldError) string {
	logrus.Debugf(fe.Tag())

	fun, ok := FiledTagFunc[fe.Tag()]
	if !ok {
		return fe.Error()
	} else {
		return fun(fe)
	}
}

func GetErrorMsgs(fes []validator.FieldError) []ErrorMsg {
	var errors []ErrorMsg
	for _, fe := range fes {
		errors = append(errors, ErrorMsg{
			Field:   fe.Field(),
			Message: GetErrorMsg(fe),
		})
	}
	return errors
}

func UniqueEmailValidator(fl validator.FieldLevel) bool {
	// email := fl.Field().String()

	// 检查email字段是否存在重复记录
	return true
}

// 初始化验证器，全局验证器
func InitValidation(r *gin.Engine) {
	logrus.Debugf("init validator")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("uniqueEmail", UniqueEmailValidator)
	}
}
