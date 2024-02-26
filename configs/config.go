package configs

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type SysConfig struct {
	Server struct {
		Host  string `yaml:"host"`
		Port  string `yaml:"port"`
		Debug bool   `yaml:"debug"`
	} `yaml:"server"`
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

var _sysConfig *SysConfig
var once sync.Once

func GetSysConfig() *SysConfig {
	// 使用 once.Do 确保初始化只执行一次
	once.Do(func() {
		// 读取YAML配置文件
		yamlFile, _ := os.ReadFile("./configs/config.yaml")
		// 解析YAML配置文件
		_ = yaml.Unmarshal(yamlFile, &_sysConfig)
	})

	return _sysConfig
}
