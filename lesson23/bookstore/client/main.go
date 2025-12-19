package main

import (
	"context"
	"lesson23/bookstore/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	// 连接server
	conn, err := grpc.NewClient("dns:///localhost:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端

	c := pb.NewBookstoreClient(conn)
	shelves, err := c.ListShelves(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("c.ListShelves failed,err:%v", err)
		return
	}
	log.Printf("shelves:%s", protojson.Format(shelves))
}
