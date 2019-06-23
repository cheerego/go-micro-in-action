package main

import (
	"context"
	"github.com/cheerego/go-micro-in-action/grpc-micro-client/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"log"
)

type Greeter struct {
}

func (s *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Print("Received grpc Hello request")
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
