package main

import (
	"context"
	"demo/proto/book"
	"demo/proto/calculate"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	TypeUrl = "type.googleapis.com/"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type calculateServer struct {
	calculate.UnimplementedCalculatorServer
}

func (*calculateServer) Add(_ context.Context, req *calculate.CalculateRequest) (*calculate.CalculateResponse, error) {
	log.Printf("c has setted: %v, c=%f", req.C != nil, req.GetC())
	var (
		student *calculate.Student = new(calculate.Student)
		user    *calculate.User    = new(calculate.User)
		data    *anypb.Any         = req.Data
		err     error
	)
	log.Printf("data: %s", data)
	log.Printf("data.TypeUrl: %s", data.TypeUrl)
	if data.TypeUrl == TypeUrl+"calculate.User" {
		err = anypb.UnmarshalTo(data, user, proto.UnmarshalOptions{})
		if err != nil {
			log.Printf("cannot unmarshal to User")
		}
		log.Printf("user = %s", protojson.Format(user))
	} else if data.TypeUrl == TypeUrl+"calculate.Student" {
		err = anypb.UnmarshalTo(data, student, proto.UnmarshalOptions{})
		if err != nil {
			log.Printf("cannot unmarshal to Student")
		}
		log.Printf("student = %s", protojson.Format(student))
	} else {
		log.Printf("unknown data type: %s", data.TypeUrl)
	}
	return &calculate.CalculateResponse{Result: req.A + req.B}, nil
}

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
	calculate.RegisterCalculatorServer(s, &calculateServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
