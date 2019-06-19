package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"google.golang.org/grpc"
	"net/http"

	hello "micro-rpc/gateway/greeter"
)

func init() {
	client.DefaultClient = client.NewClient()
	// set the default server
	server.DefaultServer = server.NewServer()
}

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
