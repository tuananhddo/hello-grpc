package hello

import (
	"context"
	"fmt"
	"hello-gr/repo"
)

type HelloServerX struct {
	UnimplementedHelloServer
}

func (h HelloServerX) GetFeature(ctx context.Context, point *Point) (*Point, error) {
	fmt.Print(point)
	fmt.Printf("Long %d \n", point.Latitude)
	fmt.Printf("Lat %d", point.Longitude)
	return point, nil
}
func (s *HelloServerX) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message: "Hello",
	}, nil
}

func (s *HelloServerX) CreateToDo(ctx context.Context, req *CreateToDoRequest) (*CreateToDoResponse, error) {
	fmt.Print(req.Content)
	repo.CreateToDoRepo()
	return &CreateToDoResponse{
		Message: "Create Success",
	}, nil
}
