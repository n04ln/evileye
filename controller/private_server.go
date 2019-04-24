package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
)

type privateServer struct {
}

func NewPrivServer() pb.PrivateServer {
	return &privateServer{}
}
