package main

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/bruce-mig/grpc-proj/proto"
)

func callSayHello(wg *sync.WaitGroup, client pb.GreetServiceClient) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%s", res.Message)

}
