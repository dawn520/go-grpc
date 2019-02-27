package service

import (
	"context"
	"grpc-server/proto"
)

type Hello struct{}

func (h *Hello) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {

	resp := new(proto.HelloReply)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}
