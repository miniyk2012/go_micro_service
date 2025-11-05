package main

import (
	"context"
	"fmt"
	"lesson20/hello_client/pb"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}

type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// streamInterceptor 客户端流式拦截器
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), nil
}

func main() {
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "miniyk.com")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed, err:%v\n", err)
		return
	}
	conn, err := grpc.NewClient("127.0.0.1:8972",
		//grpc.WithTransportCredentials(insecure.NewCredentials()), // 不安全的连接
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithStreamInterceptor(streamInterceptor),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 测试普通的 SayHello 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())

	// 测试流式 StreamHello 方法

	stream, err := c.StreamHello(context.Background())
	if err != nil {
		log.Fatalf("failed to call StreamHello: %v", err)
	}

	// 发送多个请求
	requests := []string{"first", "second", "third", "fourth", "fifth"}
	for _, req := range requests {
		if err := stream.Send(&pb.HelloRequest{Name: req}); err != nil {
			log.Fatalf("failed to send request: %v", err)
		}
		log.Printf("Sent request: %s", req)
	}

	// 关闭流并接收响应
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive reply: %v", err)
	}
	log.Printf("StreamHello reply: %s", reply.GetReply())
}
