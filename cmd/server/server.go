package main

import (
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	hello "hello-gr/proto"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7100")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := hello.HelloServerX{}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	//hello.RegisterHelloServer(gr)
	hello.RegisterHelloServer(grpcServer, &s)
	//hello.RegisterHelloServer(grpcServer, &s)
	grpcServer.Serve(lis)
}
