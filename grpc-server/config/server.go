package config

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-server/proto"
	"grpc-server/service"
	"net"
)

// 注册服务
func RegisterService(s *grpc.Server) {
	// 注册Hello Service
	proto.RegisterHelloServer(s, new(service.Hello))
}

// 启动服务
func StartServer() {

	fmt.Println("start server")

	var host = SERVER_HOST
	var port = SERVER_PORT
	if Cfg.ServerHost != "" {
		host = Cfg.ServerHost
	}
	if Cfg.ServerPort != "" {
		port = Cfg.ServerPort
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