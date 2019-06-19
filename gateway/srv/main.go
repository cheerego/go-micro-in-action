package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	"micro-rpc/gateway/greeter"
)

type Greeter struct{}

func (s *Greeter) Say(context.Context, *greeter.Request, *greeter.Response) error {
	panic("implement me")
}

func (s *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	greeter.RegisterSayHandler(service.Server(), new(Greeter))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
