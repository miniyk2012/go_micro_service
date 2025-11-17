package main

import (
	"context"
	"fmt"
	"io"
	"lesson20/hello_server/pb"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Reply: "hello world",
	}, nil
}

func (s *server) StreamHello(stream grpc.ClientStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	// 持续接收request, 计算一共有几个req, 然后再返回中HelloResponse.Reply中返回数量
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv failed: %v", err)
			return status.Errorf(codes.Unknown, "stream.Recv failed: %v", err)
		}
		log.Printf("Receive a message: %s", req.GetName())
		count++
	}
	return stream.SendAndClose(&pb.HelloResponse{
		Reply: "Total requests: " + strconv.Itoa(count),
	})

}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %s\n", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// valid 校验认证信息.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token认证的逻辑
	// 这里是为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}

// unaryInterceptor 服务端一元拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	fmt.Printf("req's type is %T\n", req) // req's type is *pb.HelloRequest
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	m, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	fmt.Printf("m's type is %T\n", m) // m's type is *pb.HelloRespons
	return m, err
}

// streamInterceptor 服务端流拦截器
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		log.Printf("null md in streaming, stop handle")
		return status.Errorf(codes.InvalidArgument, "missing metadata in streaming")
	}
	if !valid(md["authorization"]) {
		log.Printf("invalid token in streaming, stop handle")
		return status.Errorf(codes.Unauthenticated, "invalid token in streaming")
	}

	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		fmt.Printf("RPC failed with error %v\n", err)
	}
	return err
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	// 加载证书信息
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed, err:%v\n", err)
		return
	}
	s := grpc.NewServer(grpc.Creds(creds),
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	) // 创建grpc服务
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
