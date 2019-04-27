package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewPublicClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{ScreenName: "shinka", Password: "morisama"})
	if err != nil {
		log.Fatalf("cannnot login: %v", err)
	}

	log.Println(r)
	newCtx := metadata.AppendToOutgoingContext(ctx, "Authorization", "Bearer "+r.Token)
	cp := pb.NewPrivateClient(conn)
	u, err := cp.GetUserInfo(newCtx, &pb.UserInfoReq{UserName: 1})

	if err != nil {
		log.Fatalf("cannot get user : %v\n", err)
	}
	log.Println(u)
}
