package main

import (
	"context"
	"github.com/kfoozminus/hello-grpc/cmd/proxy"
	"log"
	"net"

	"github.com/kfoozminus/hello-grpc/build/gen"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServiceServer struct{}

func (c *helloServiceServer) Intro(ctx context.Context, in *foo.IntroRequest) (*foo.IntroResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &foo.IntroResponse{Msg: "Hello " + in.Name}, nil
}

func (c *helloServiceServer) IntroAgain(ctx context.Context, in *foo.IntroRequest) (*foo.IntroResponse, error) {
	log.Printf("Received Again: %v", in.Name)
	return &foo.IntroResponse{Msg: "Hello Again, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	foo.RegisterHelloServiceServer(s, &helloServiceServer{})

	// Register reflection service on gRPC server.
	//reflection.Register(s)

	go proxy.Call()

	log.Println("Server running at port :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
