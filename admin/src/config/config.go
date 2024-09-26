package config

import (
	"encoding/json"
	"fmt"
	logDef "log"
	"os"

	"siteol.com/smart/src/common/log"
)

var JsonConfig *Config

// Config 配置对象
type Config struct {
	Server Server // 服务配置
	MySQL  MySQL  // MySQL数据配置
	Redis  Redis  // Redis配置
}

// Server 服务配置
type Server struct {
	Name     string // 服务名
	Port     string // 服务端口
	LogLevel string // 日志等级
	LogRoot  int    // 日志展示路径范围
	Debug    bool   // 调试模式，打开时部分处理跳过或进入Mock处理
}

// MySQL MySQL数据配置组
type MySQL struct {
	Plat string // 授权平台
	Blog string // 博客应用
}

// Redis Redis配
type Redis struct {
	Addr     string // Redis服务地址
	DB       int    // 数据库号
	Password string // 密码（可为空）
}

// 配置初始化
func init() {
	SetEnv("server")                          // 初始化环境变量
	ReadConfig()                              // 读取
	log.SetLevel(JsonConfig.Server.LogLevel)  // 日志等级
	log.SetAutoSplit(log.CronDaily)           // 日志切割
	log.SetAppRoot(JsonConfig.Server.LogRoot) // 日志深度(2表示上级../logs/）
}

// ReadConfig 启动读取配置文件
func ReadConfig() {
	JsonConfig = &Config{}
	filePath := fmt.Sprintf("%s/config/config_%s.json", WorkDir, SysEnv)
	file, _ := os.Open(filePath)
	defer func() { _ = file.Close() }()
	decoder := json.NewDecoder(file)
	for decoder.More() {
		err := decoder.Decode(JsonConfig)
		if err != nil {
			logDef.Printf("Read %s Error: %v \n", filePath, err)
			os.Exit(1)
		}
	}
	logDef.Printf("Config Read %s Success .\n", filePath)
}
