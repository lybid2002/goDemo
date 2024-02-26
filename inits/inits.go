package inits

import (
	"gotest/db"
	"gotest/logs"
)

func InitOpration() {
	logs.InitConfig()

	// 测试 db 模块中的 init 方法和连接数据库测试
	db.Test()
}
