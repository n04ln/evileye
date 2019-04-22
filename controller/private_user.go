package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

type PrivateUserHandler struct {
}

func (puh *PrivateUserHandler) GetUserInfo(c context.Context, uinforeq *pb.UserInfoReq) (*pb.User, error) {
	panic("not impl")
}

func (puh *PrivateUserHandler) GetUserList(c context.Context, userlistreq *pb.GetUserListReq) (*pb.User, error) {
	panic("not impl")
}
