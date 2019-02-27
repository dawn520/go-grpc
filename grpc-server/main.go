package main

import (
	"grpc-server/config"
)

func main() {

	// 初始化配置
	config.Run()
	// 启动服务
	config.StartServer()
}
