package main

import (
	"context"
	"fmt"
	"lesson45/pb"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// gRPC test

// 编写一个gRPC客户端，测试我们的gRPC Server是否正常

// 使用bufconn构建测试链接，避免使用实际端口号启动服务

const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	gs := NewGRPCServer(addService{})
	pb.RegisterAddServer(s, gs)
	go func() {
		if err := s.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

func TestServer(t *testing.T) {

	conn, err := grpc.NewClient(
		"passthrough:///",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufDialer),
	)
	if err != nil {
		t.Fail()
	}
	defer conn.Close()

	c := pb.NewAddClient(conn)
	resp, err := c.Sum(context.Background(), &pb.SumRequest{A: 1, B: 2})
	assert.Nil(t, err)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, resp)
	assert.Equal(t, int64(3), resp.V)

	resp2, err2 := c.Concat(context.Background(), &pb.ConcatRequest{A: "hello", B: "world"})
	assert.Nil(t, err2)
	assert.NotNil(t, resp2)
	assert.Equal(t, "helloworld", resp2.V)
}
