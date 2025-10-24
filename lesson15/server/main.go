package main

import (
	"flag"
	"fmt"
	"io"
	"lesson15/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type routeGuideServer struct {
	pb.UnimplementedGreetingServiceServer
}

func (*routeGuideServer) SayMultiple(stream grpc.ClientStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	reply := "你好："
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Msg: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += req.GetMsg() + ", "
	}
}
func (*routeGuideServer) RevMultiple(req *pb.HelloRequest, stream grpc.ServerStreamingServer[pb.HelloResponse]) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _, word := range words {
		reply := &pb.HelloResponse{
			Msg: word + ":" + req.GetMsg(),
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
	return nil
}

func (*routeGuideServer) Chat(stream grpc.BidiStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		reply := &pb.HelloResponse{
			Msg: "hello:" + req.GetMsg(),
		}
		if err = stream.Send(reply); err != nil {
			return err
		}
	}
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{}
	return s
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreetingServiceServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
