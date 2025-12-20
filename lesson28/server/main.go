package main

import (
	"context"
	"flag"
	"fmt"
	"lesson28/pb"
	"net"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 8972, "服务端口")

type server struct {
	pb.UnimplementedGreeterServer
	Addr string
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Reply: fmt.Sprintf("hello %s, addr=%s", req.Name, s.Addr),
	}, nil
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// 启动服务
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	s := grpc.NewServer() // 创建grpc服务
	pb.RegisterGreeterServer(s, &server{Addr: addr})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve, err:%v\n", err)
		return
	}
}
