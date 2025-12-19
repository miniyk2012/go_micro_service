package main

import (
	"context"
	"fmt"
	"lesson23/bookstore/pb"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
	l, err := net.Listen("tcp", ":8972")
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

	// grpc-Gateway
	conn, err := grpc.NewClient(
		"127.0.0.1:8972",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("grpc conn failed, err:%v\n", err)
		return
	}

	gwmux := runtime.NewServeMux()
	pb.RegisterBookstoreHandler(context.Background(), gwmux, conn)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	fmt.Println("grpc-Gateway serve on :8090...")
	err = gwServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
