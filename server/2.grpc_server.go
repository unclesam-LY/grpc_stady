package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc_stady/grpc_proto/hello_grpc"
	"net"
)

type HelloService struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (HelloService) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	return &hello_grpc.HelloResponse{
		Name:    "sam",
		Message: "fuck",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		grpclog.Fatal(err)
		return
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	servers := HelloService{}
	// 注册服务
	hello_grpc.RegisterHelloServiceServer(s, servers)

	// 启动服务
	err = s.Serve(listen)
	if err != nil {
		grpclog.Fatal(err)
	}
}
