package hello

import (
	"context"
	"fmt"
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
