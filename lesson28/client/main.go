package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "lesson28/pb"
)

func main() {
	// 指定连接server
	//conn, err := grpc.NewClient("127.0.0.1:8972",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	// dns解析
	//conn, err := grpc.NewClient("dns:///localhost:8972",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	// 自定义解析
	conn, err := grpc.NewClient("yk:///resolver.yangkai.com",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithResolvers(&ykResolverBuilder{}),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 这里设置初始策略
	)
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "yk"})
		if err != nil {
			fmt.Printf("c.SayHello failed, err:%v\n", err)
			return
		}
		// 拿到了RPC响应
		fmt.Printf("resp:%v\n", resp.GetReply())
	}
}
