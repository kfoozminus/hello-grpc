package proxy

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kfoozminus/hello-grpc/build/gen"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:8080", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := foo.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":4040", mux)
}

func Call() {
	flag.Parse()

	log.Println("Proxy server running at port 4040")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}