package main

import (
	"context"
	"io"
	"log"
	"sync"

	pb "github.com/bruce-mig/grpc-proj/proto"
)

func callSayHelloServerStream(wg *sync.WaitGroup, client pb.GreetServiceClient, names *pb.NamesList) {
	defer wg.Done()
	log.Print("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
