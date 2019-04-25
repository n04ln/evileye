package controller

import (
	"context"

	"github.com/NoahOrberg/evileye/config"
	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pth *privateServer) Tarekomi(c context.Context, tarekomireq *pb.TarekomiReq) (*pb.Empty, error) {
	nt := entity.Tarekomi{
		Status:       0,
		Threshold:    config.GetConfig().Threshold,
		TargetUserID: tarekomireq.Tarekomi.TargetUserId,
		URL:          tarekomireq.Tarekomi.Url,
		Description:  tarekomireq.Tarekomi.Desc,
	}

	_, err := pth.TUsecase.Store(c, nt)
	if err != nil {
		log.L().Error("Tarekomi Store failed", zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "Database down")
	}

	return &pb.Empty{}, nil
}

func (pth *privateServer) TarekomiBoard(c context.Context, tbordreq *pb.TarekomiBoardReq) (*pb.TarekomiBoardRes, error) {
	tb, err := pth.TUsecase.GetTarekomiBoard(c, tbordreq.Limit, tbordreq.Offset)
	if err != nil {
		log.L().Error("GetTarekomiBoard failed", zap.Error(err))
		return &pb.TarekomiBoardRes{}, status.Error(codes.Internal, "Database down")
	}

	rtb := &pb.TarekomiBoardRes{Tarekomis: tb.Tarekomis}

	return rtb, nil
}

func (pth *privateServer) AddStar(c context.Context, addstarreq *pb.AddStarReq) (*pb.Empty, error) {
	panic("not impl")
	// TODO: starを実装したら作る
}

func (pth *privateServer) GetStaredTarekomi(c context.Context, e *pb.Empty) (*pb.TarekomiSummaries, error) {
	panic("not impl")
	// TODO: starを実装したら作る
}
