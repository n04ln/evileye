package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

type PrivateUserHandler struct {
}

func (puh *PrivateUserHandler) GetUserInfo(c context.Context, uinforeq *pb.UserInfoReq) (*pb.User, error) {
	return nil, nil
}
