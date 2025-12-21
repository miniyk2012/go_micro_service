package main

import (
	"context"
	"flag"
	"fmt"
	"lesson35/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	serviceName = "hello"
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
	ipinfo, err := GetOutboundIP()
	if err != nil {
		fmt.Printf("get outbound ip failed, err:%v\n", err)
		return
	}

	// 启动服务
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	addr := fmt.Sprintf("%s:%d", ipinfo.String(), *port)
	fmt.Printf("server start at %s\n", addr)
	s := grpc.NewServer() // 创建grpc服务
	pb.RegisterGreeterServer(s, &server{Addr: addr})
	healthpb.RegisterHealthServer(s, health.NewServer()) // consul 发来健康检查的RPC请求，这个负责返回OK

	// 注册consul
	consul, err := NewConsul()
	if err != nil {
		fmt.Printf("NewConsul failed, err:%v\n", err)
		return
	}
	err = consul.RegisterService(serviceName, ipinfo.String(), *port)
	if err != nil {
		fmt.Printf("RegisterService failed, err:%v\n", err)
		return
	}
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve, err:%v\n", err)
		return
	}
}
