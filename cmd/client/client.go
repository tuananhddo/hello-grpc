package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "hello-gr/proto"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:7100", opts...)
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Print(conn)

	client := hello.NewHelloClient(conn)
	ctx := context.Background()

	p := hello.Point{
		Latitude:  0,
		Longitude: 0,
	}
	data, errs := client.GetFeature(ctx, &p)
	if errs != nil {
		fmt.Print(errs)
	}
	if data == nil {
		fmt.Print("NIL")
	}
	fmt.Print(data.Longitude)
	defer conn.Close()
	fmt.Print("END")

}
