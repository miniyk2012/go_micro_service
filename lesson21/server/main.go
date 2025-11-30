package main

import (
	"context"
	"flag"
	helloworldpb "lesson21/proto"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{Reply: in.Name + " world"}, nil
}

func main() {
	flag.Parse()
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	helloworldpb.RegisterGreeterServer(s, NewServer())
	// 新增：注册反射 GPRC API
	reflection.Register(s) // 让服务支持反射
	// 启动gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		grpclog.Fatal(s.Serve(lis))
	}()

	if err := runHttpProxy(); err != nil {
		grpclog.Fatal(err)
	}
}

func runHttpProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := helloworldpb.RegisterGreeterHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("Serving http on 0.0.0.0:8082")
	return http.ListenAndServe(":8082", mux)
}
