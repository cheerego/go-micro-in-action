package server

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/cheerego/go-micro-in-action/grpc-gateway/proto"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:58863", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":9090", mux)
}

func main() {
	flag.Parse()

	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
