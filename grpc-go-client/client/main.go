package main

import (
	pb "github.com/cheerego/go-micro-in-action/grpc-go-client/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:54686", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &pb.Request{Name: "hkn"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Greeting)
	defer conn.Close()
}
