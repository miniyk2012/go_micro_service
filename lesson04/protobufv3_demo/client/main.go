package main

import (
	"context"
	"flag"
	"log"
	"protobufv3_demo/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	defaultQuery = "world"
)

var (
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
	query = flag.String("query", defaultQuery, "Query to search")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearchServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Search(ctx, &pb.SearchRequest{Query: *query})
	if err != nil {
		log.Fatalf("could not search: %v", err)
	}
	v, _ := protojson.Marshal(r.GetResult())
	log.Printf("Searching: %s", v)
}
