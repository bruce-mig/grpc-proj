package main

import (
	"log"
	"sync"

	pb "github.com/bruce-mig/grpc-proj/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

var wg *sync.WaitGroup

func main() {
	conn, err := grpc.Dial("0.0.0.0"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Alice", "Bob", "Bruce"},
	}

	wg = &sync.WaitGroup{}

	// wg.Add(1)
	// go callSayHello(wg, client)
	// wg.Add(1)
	// go callSayHelloServerStream(wg, client, names)
	// wg.Add(1)
	// go callSayHelloClientStream(wg, client, names)
	wg.Add(1)
	go callHelloBidirectionalStream(wg, client, names)
	wg.Wait()
}
