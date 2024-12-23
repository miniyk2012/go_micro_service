package main

import (
	"context"
	"fmt"
	"hello_server/pb"
	"net"

	"google.golang.org/grpc"
)



type helloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Reply: "Hello " + in.GetName(),
	}, nil
}

type addServer struct {
	pb.UnimplementedAddServiceServer
}

func (s *addServer) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Result: in.A + in.B,
    Code: 0,
    Msg: "success",
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	pb.RegisterGreeterServer(s, &helloServer{})
	pb.RegisterAddServiceServer(s, &addServer{})
  // 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}

