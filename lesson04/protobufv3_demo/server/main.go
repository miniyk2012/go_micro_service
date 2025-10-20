package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"protobufv3_demo/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSearchServiceServer
}

func (s *server) Search(_ context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	reply := req.Query + " - response from server"
	// Implement your search logic here
	return &pb.SearchResponse{
		Result: &pb.Result{Title: reply, Week: pb.Week_MONDAY},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %f", err)
	}
}
