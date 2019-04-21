package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

type PublicAccountStateHandler struct {
}

func (pas *PublicAccountStateHandler) Login(c context.Context, loginreq *pb.LoginRequest) (*pb.LoginRes, error) {
	return nil, nil
}
