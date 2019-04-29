package controller

import (
	"context"

	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (puh *privateServer) GetUserInfo(c context.Context, uinforeq *pb.UserInfoReq) (*pb.User, error) {
	ui := interceptor.GetUserMetaData(c)

	u, err := puh.URepository.UserGetByID(c, ui.ID)
	if err != nil {
		log.L().Error("UserGetByID error", zap.Error(err))
		return &pb.User{}, status.Error(codes.Internal, "Database down")
	}

	return &pb.User{
		UserId:   u.ID,
		UserName: u.ScreenName,
	}, nil
}

func (puh *privateServer) GetUserList(c context.Context, userlistreq *pb.GetUserListReq) (*pb.Users, error) {
	us, err := puh.URepository.UserGetByIDList(c, userlistreq.Limit, userlistreq.Offset)
	if err != nil {
		log.L().Error("repository error",
			zap.Int64("limit", userlistreq.GetLimit()),
			zap.Int64("offset", userlistreq.GetOffset()),
			zap.Error(err))
		return &pb.Users{}, status.Error(codes.Internal, "Database down")
	}

	usr := make([]*pb.User, 0, len(us))
	for _, eu := range us {
		u := &pb.User{
			UserId:   eu.ID,
			UserName: eu.ScreenName,
		}

		tarekomi, err := puh.TRepository.GetTarekomiApproved(c, u.UserId)
		if err != nil {
			log.L().Error("GetTarekomiApproved error", zap.Error(err))
		}
		u.Tarekomis = tarekomi
		usr = append(usr, u)
	}

	return &pb.Users{
		Users: usr,
	}, nil
}
