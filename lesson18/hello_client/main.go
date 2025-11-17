package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"code.xxx.com/backend/hello_client/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	// "google.golang.org/grpc/metadata"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "miniyk", "通过-name告诉server你是谁")

func handelError(err error) {
	s := status.Convert(err)
	for _, d := range s.Details() {
		switch info := d.(type) {
		case *errdetails.QuotaFailure:
			fmt.Printf("Quota failure: %s\n", info)
		default:
			fmt.Printf("Unexpected type: %#v\n", info)
		}
	}
}

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	// 拿到了RPC响应
	log.Printf("resp:%v\n", resp.GetReply())
	resp, err = c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err == nil {
		log.Fatal("c.SayHello twice should failed")
	}
	handelError(err)
}
