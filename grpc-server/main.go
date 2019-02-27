package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-server/config"
	"grpc-server/proto"
	"grpc-server/service"
	"net"
)

func main() {

	// 初始化配置
	config.Run()
	// 启动服务
	startServer()
}

func startServer() {

	fmt.Println("start server")

	var cfg = config.Cfg
	var host = config.SERVER_HOST
	var port = config.SERVER_PORT
	if cfg.ServerHost != "" {
		host = cfg.ServerHost
	}
	if cfg.ServerPort != "" {
		port = cfg.ServerPort
	}
	address := host + ":" + port
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册Service
	RegisterService(s)

	fmt.Println("Listen on " + address)

	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}


func RegisterService(s *grpc.Server) {
	// 注册Hello Service
	proto.RegisterHelloServer(s, &service.Hello{})
}
