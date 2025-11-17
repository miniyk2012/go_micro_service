package main

import (
	"context"
	"flag"
	"lesson21/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "miniyk", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := proto.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发起普通RPC调用
	resp, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Printf("first c.SayHello failed, err:%v", err)
		return
	}
	log.Printf("c.SayHello success, resp:%v", resp)
}
