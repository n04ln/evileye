package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

func (puh *privateServer) GetUserInfo(c context.Context, uinforeq *pb.UserInfoReq) (*pb.User, error) {
	panic("not impl")
}

func (puh *privateServer) GetUserList(c context.Context, userlistreq *pb.GetUserListReq) (*pb.Users, error) {
	panic("not impl")
}
