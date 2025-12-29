package main

import (
	"flag"
	"fmt"
	"lesson47/pb"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	httpAddr = flag.Int("http-addr", 8080, "HTTP端口")
	gRPCAddr = flag.Int("grpc-addr", 8972, "gRPC端口")
)

func main() {
	srv := NewService()
	var g errgroup.Group
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		server := NewHTTPServer(srv)
		return http.Serve(httpListener, server)
	})

	g.Go(func() error {
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *gRPCAddr, err)
			return err
		}
		defer grpcListener.Close()

		s := grpc.NewServer()
		pb.RegisterAddServer(s, NewGRPCServer(srv))
		return s.Serve(grpcListener)
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("g.Wait failed, err:%v\n", err)
	}
}
