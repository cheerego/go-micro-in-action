package main

import (
	"context"
	"fmt"
	"github.com/cheerego/go-micro-in-action/grpc-micro-client/proto"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := greeter.NewGreeterService("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &greeter.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
