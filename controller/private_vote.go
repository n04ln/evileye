package controller

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pvh *privateServer) Vote(c context.Context, votereq *pb.VoteReq) (*pb.Empty, error) {
	ui := interceptor.GetUserMetaData(c)

	nv := &entity.Vote{
		UserID:      ui.ID,
		TarekomiID:  votereq.TarekomiId,
		Description: votereq.Desc,
	}

	err := pvh.VRepository.NewVoting(c, nv, pvh.IClient)
	if err != nil {
		log.L().Error("NewVoting failed", zap.Any("req", votereq), zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "Database down")
	}

	return &pb.Empty{}, nil
}
