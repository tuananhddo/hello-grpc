package main

import (
	"google.golang.org/grpc"
	hello "hello-gr/proto"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7100")
	var opts []grpc.ServerOption
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := hello.HelloServerX{}
	grpcServer := grpc.NewServer(opts...)
	//hello.RegisterHelloServer(gr)
	hello.RegisterHelloServer(grpcServer, s)
	grpcServer.Serve(lis)
}
