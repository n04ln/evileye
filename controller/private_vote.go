package controller

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/interceptor"
	pb "github.com/NoahOrberg/evileye/protobuf"
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

	err := pvh.VRepository.NewVoting(c, nv)
	if err != nil {
		return &pb.Empty{}, status.Error(codes.Internal, "Database down")
	}

	return &pb.Empty{}, nil
}
