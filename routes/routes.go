package routes

import (
	"gotest/controllers"
	"gotest/middlewares"
	"gotest/models"

	"github.com/gin-gonic/gin"
)

// * **全局应用**：
// go`router := gin.Default()
// router.Use(LoggingMiddleware())`

// * **应用到路由组**：
// go`apiGroup := router.Group("/api")
// apiGroup.Use(LoggingMiddleware())`

// * **应用到单个路由**：
// go`router.GET("/user/:id", LoggingMiddleware(), func(c *gin.Context) {
// })`

func InitRoutes() *gin.Engine {
	router := gin.Default()

	models.InitValidation(router)

	router.Use(middlewares.LoggerMiddleware())
	// router.Use(middlewares.ValidatorMiddleware)

	router.Use(middlewares.ErrorHandler)

	// 默认路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// 用户路由组
	userGroup := router.Group("/api/user")
	{
		userGroup.POST("/register", controllers.GetUserController().Register)
		userGroup.POST("/login", controllers.GetUserController().Login)
		userGroup.GET("/register", controllers.GetUserController().Register)
		userGroup.GET("/login", controllers.GetUserController().Login)
	}

	return router
}
