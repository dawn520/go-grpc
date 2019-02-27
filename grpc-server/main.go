package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-server/proto"
	"grpc-server/service"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func RegisterService(s *grpc.Server) {
	// 注册Hello Service
	proto.RegisterHelloServer(s, &service.Hello{})
}

func main() {

	fmt.Println("start server")
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册Service
	RegisterService(s)

	fmt.Println("Listen on " + Address)

	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
