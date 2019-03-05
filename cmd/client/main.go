package main

import (
	"context"
	"github.com/kfoozminus/hello-grpc/build/gen"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	defaultName = "Anonnymous"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client couldn't be established: %v", err)
	}
	defer conn.Close()

	client := foo.NewHelloServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Intro(ctx, &foo.IntroRequest{Name: name})
	if err != nil {
		log.Fatalf("could not agree: %v", err)
	}
	log.Printf("Greeting: %s", r.Msg)

	r, err = client.IntroAgain(ctx, &foo.IntroRequest{Name: name})
	if err != nil {
		log.Fatalf("could not agree: %v", err)
	}
	log.Printf("Greeting: %s", r.Msg)
}