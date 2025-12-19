package main

import (
	"context"
	"fmt"
	"lesson23/bookstore/pb"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var customFunc recovery.RecoveryHandlerFunc

func main() {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	// 创建server
	srv := server{
		bs: &bookstore{db: db},
	}
	// 启动gRPC服务
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	// Define customfunc to handle panic
	customFunc = func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []recovery.Option{
		recovery.WithRecoveryHandler(customFunc),
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(
		recovery.UnaryServerInterceptor(opts...),
	))
	// 注册服务
	pb.RegisterBookstoreServer(s, &srv)
	go func() {
		fmt.Println(s.Serve(l))
	}()

	// 同一个端口分别处理gRPC和HTTP
	// 1. 创建gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterBookstoreHandlerFromEndpoint(
		context.Background(), gwmux, "127.0.0.1:8090", dops); err != nil {
		log.Fatalf("RegisterBookstoreHandlerFromEndpoint failed, err:%v\n", err)
		return
	}

	// 2. 新建HTTP mux
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 3. 定义HTTP Server
	gwServer := &http.Server{
		Addr:    "127.0.0.1:8090",
		Handler: grpcHandlerFunc(s, mux),
	}

	// 4. 启动服务
	fmt.Println("serving on 127.0.0.1:8090...")
	gwServer.Serve(l)
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
