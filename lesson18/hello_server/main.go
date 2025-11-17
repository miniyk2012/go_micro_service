package main

import (
	"context"
	"fmt"
	"hello_server/proto"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

// grpc server

type server struct {
	proto.UnimplementedGreeterServer
	mu    sync.Mutex
	count map[string]int // 记录每个name的请求次数
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(_ context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.GetName()]++
	if s.count[in.GetName()] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "限制每个name调用一次",
				}},
			}, in)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	reply := "hello " + in.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	proto.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
