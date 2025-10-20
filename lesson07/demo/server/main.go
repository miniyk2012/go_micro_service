package main

import (
	"context"
	"flag"
	"demo/proto/book"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	book.UnimplementedBookServiceServer
}

func (*server) CreateBook(c context.Context, b *book.Book) (*book.Book, error) {
	log.Printf("receive a book create request: %s", protojson.Format(b))
	b.Name += " - response from server"
	b.Price.MarketPrice += 100
	b.Price.SalePrice += 200
	return b, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	book.RegisterBookServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
