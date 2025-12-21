package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"lesson35/pb"
)

const (
	serviceName = "hello"
)

// grpc 客户端
// 调用server端的 SayHello 方法
var name = flag.String("name", "yk", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 1. 连接consul
	cc, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}
	// 2. 根据服务名称查询服务实例
	// cc.Agent().Services()  // 列出所有的
	serviceMap, err := cc.Agent().ServicesWithFilter(fmt.Sprintf("Service==`%s`", serviceName)) // 查询服务名称是hello的所有服务节点
	if err != nil {
		fmt.Printf("query `hello` service failed, err:%v\n", err)
		return
	}
	var addr string
	for k, v := range serviceMap {
		fmt.Printf("%s:%#v\n", k, v)
		addr = fmt.Sprintf("%s:%d", v.Address, v.Port) // 取第一个机器的address和port
	}
	// 3. 随机取serviceMap中的某个value:
	for _, v := range serviceMap {
		addr = fmt.Sprintf("%s:%d", v.Address, v.Port)
		break
	}
	// 连接server
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	// 拿到了RPC响应
	fmt.Printf("resp:%v\n", resp.GetReply())
}
