package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	hello "hello-gr/proto"
	"net/http"
)

func main() {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	ctx := context.Background()
	err := hello.RegisterHelloHandlerFromEndpoint(ctx, mux, ":7100", opts)
	err = hello.RegisterHelloHandlerFromEndpoint(ctx, mux, ":7100", opts)

	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":7200", mux)
	if err != nil {
		panic(err)
	}

}
