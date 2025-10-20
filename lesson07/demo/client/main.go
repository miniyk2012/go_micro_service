package main

import (
	"context"
	"demo/proto/author"
	"demo/proto/book"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := book.NewBookServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateBook(ctx, &book.Book{
		Name:   "Go Programming Language",
		Price:  &book.Price{MarketPrice: 100, SalePrice: 200},
		Author: &author.Info{Name: "yangkai"},
		Date:   timestamppb.New(time.Date(2025, 6, 6, 6, 16, 0, 0, time.Local)),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("create one book: %s", protojson.Format(r))
}
