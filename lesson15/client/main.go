package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"lesson15/pb"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func sayMultipleDemo(client pb.GreetingServiceClient) {

	var names = []string{"gopher", "golang", "grpc"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.SayMultiple(ctx)
	if err != nil {
		log.Fatalf("client.SayMultiple failed: %v", err)
	}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Msg: name}); err != nil {
			log.Fatalf("client.SayMultiple failed: %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.SayMultiple failed: %v", err)
	}
	log.Printf("SayMultiple Reply: %s", reply.Msg)
}

func revMultipleDemo(client pb.GreetingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RevMultiple(ctx, &pb.HelloRequest{Msg: "hello"})
	if err != nil {
		log.Fatalf("client.RevMultiple failed: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream.Recv failed: %v", err)
		}
		log.Printf("RevMultiple Msg: %s", resp.Msg)
	}
}

func chatDemo(client pb.GreetingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("client.Chat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.Chat stream.Recv failed: %v", err)
			}
			log.Printf("Got message %s", in.Msg)
		}
	}()
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if cmd == "exit" {
			break
		}
		if err := stream.Send(&pb.HelloRequest{Msg: cmd}); err != nil {
			log.Fatalf("client.Chat stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetingServiceClient(conn)
	revMultipleDemo(client)
	fmt.Println()
	sayMultipleDemo(client)
	fmt.Println()
	chatDemo(client)
}
