package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewPublicClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.HealthCheck(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Commit Hash: %s \n", r.CommitHash)
	log.Printf("Build At: %s \n", r.BuildTime)
}
