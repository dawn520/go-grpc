package config

import "github.com/zctod/tool/config"

type configure struct {
	ServerHost  string `config:"default:127.0.0.1"`
	ServerPort  string `config:"default:50052"`
	SqlDriver   string `config:"default:postgres"`
	SqlHost     string `config:"default:127.0.0.1"`
	SqlPort     string `config:"default:5432"`
	SqlUsername string `config:"default:root"`
	SqlPassword string `config:"default:root"`
	SqlDb       string `config:"default:test"`
}

var Cfg = &configure{}

// 初始化配置
func Run() {
	// 初始化配置文件
	config.InitConfig(Cfg, PATH_ENV)
}
