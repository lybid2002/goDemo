package main

// go install github.com/pilu/fresh@latest
// 安装 fresh，控制台执行 fresh，直接启动，自动热加载

// import "project/routes" 这样的语句时，project 通常是一个模块名，而不是项目目录的根路径名。这里的 project 在 go.mod 文件中定义的模块名

import (
	"encoding/json"
	"fmt"
	"gotest/configs"
	"gotest/inits"
	"gotest/routes"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	inits.InitOpration()

	cfg := configs.GetSysConfig()
	jsonBytes, _ := json.MarshalIndent(cfg, "", "  ")

	fmt.Printf("Fmt Config Info: \n%s\n", string(jsonBytes))
	logrus.Debugf("Debug Config Info: \n%s\n", string(jsonBytes))
	logrus.Infof("Info Config Info: \n%s\n", string(jsonBytes))
	logrus.Warnf("Warn Config Info: \n%s\n", string(jsonBytes))
	logrus.Errorf("Error Config Info: \n%s\n", string(jsonBytes))

	// 初始化路由
	router := routes.InitRoutes()

	// 运行服务器
	fmt.Printf("Starting server on port %s\n", cfg.Server.Port)

	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Error running server: %s \n", err)
	}
}
