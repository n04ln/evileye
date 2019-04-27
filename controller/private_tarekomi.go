package controller

import (
	"context"

	"github.com/NoahOrberg/evileye/config"
	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/wayback"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (pth *privateServer) Tarekomi(c context.Context, tarekomireq *pb.TarekomiReq) (*pb.Empty, error) {
	u, err := pth.URepository.UserGetByName(c, tarekomireq.Tarekomi.TargetUserName)
	if err != nil {
		log.L().Error("cannot get requested user name : ", zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "cannot find requested username")
	}

	wayurl, err := wayback.NewTarekomi(tarekomireq.Tarekomi.Url)
	if err != nil {
		log.L().Error("cannot save requested url : ", zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "cannot save request url")
	}

	nt := entity.Tarekomi{
		Status:       0,
		Threshold:    config.GetConfig().Threshold,
		TargetUserID: u.ID,
		URL:          wayurl,
		Description:  tarekomireq.Tarekomi.Desc,
	}

	_, err = pth.TRepository.Store(c, nt)
	if err != nil {
		log.L().Error("Tarekomi Store failed", zap.Error(err))
		return &pb.Empty{}, status.Error(codes.Internal, "Database down")
	}

	return &pb.Empty{}, nil
}

func (pth *privateServer) TarekomiBoard(c context.Context, tbordreq *pb.TarekomiBoardReq) (*pb.TarekomiBoardRes, error) {
	tb, err := pth.TRepository.GetTarekomiBoard(c, tbordreq.Limit, tbordreq.Offset)
	if err != nil {
		log.L().Error("GetTarekomiBoard failed", zap.Error(err))
		return &pb.TarekomiBoardRes{}, status.Error(codes.Internal, "Database down")
	}

	rtb := &pb.TarekomiBoardRes{Tarekomis: tb.Tarekomis}

	return rtb, nil
}

func (pth *privateServer) AddStar(c context.Context, addstarreq *pb.AddStarReq) (*pb.Empty, error) {
	ui := interceptor.GetUserMetaData(c)

	ns := &entity.Star{
		UserID:     ui.ID,
		TarekomiID: addstarreq.TarekomiId,
	}

	_, err := pth.SRepository.Store(c, ns)
	if err != nil {
		return &pb.Empty{}, status.Error(codes.Internal, "Database Down")
	}

	return &pb.Empty{}, nil
}

func (pth *privateServer) GetStaredTarekomi(c context.Context, e *pb.Empty) (*pb.TarekomiSummaries, error) {
	ui := interceptor.GetUserMetaData(c)

	tids, err := pth.SRepository.GetStaredTarekomiID(c, ui.ID)
	if err != nil {
		return &pb.TarekomiSummaries{}, status.Error(codes.Internal, "Databae Down")
	}

	ts := new(pb.TarekomiSummaries)

	for _, tid := range tids {
		t, err := pth.TRepository.GetTarekomiFromID(c, tid)
		if err != nil {
			return ts, status.Error(codes.Internal, "cannot get requested tarekomi")
		}
		ts.Tarekomis = append(ts.Tarekomis, &t)
	}

	return ts, nil
}
