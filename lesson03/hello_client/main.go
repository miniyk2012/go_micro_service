package main

import (
	"flag"
	"log"
	"time"

	"com.xx.yangkai/hello_client/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var a = flag.Int("a", 0, "a")
var b = flag.Int("b", 0, "b")

// go run main.go -a=-2039 -b=23
func main() {
	flag.Parse()
	// 连接server
	conn, err := grpc.NewClient("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建客户端
	c := proto.NewAddServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
	defer cancel()

	resp, err := c.Add(ctx, &proto.AddRequest{
		A: int32(*a),
		B: int32(*b),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("result: %d", resp.GetResult())
	log.Printf("resp: %v", resp)
}