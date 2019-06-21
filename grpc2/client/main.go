package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "micro-rpc/grpc2/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:54686", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}
	c := pb.NewSayClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &pb.Request{Name: "hkn"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Msg)
	defer conn.Close()
}
