package config

import "github.com/zctod/tool/config"

type configure struct {
	ServerHost string `config:"default:127.0.0.1"`
	ServerPort string `config:"default:50052"`
}

var Cfg = &configure{}

// 初始化配置
func Run() {
	// 初始化配置文件
	config.InitConfig(Cfg, PATH_ENV)
}
