package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

type PrivateTarekomiHandler struct {
}

func (pth *PrivateTarekomiHandler) Tarekomi(c context.Context, tarekomireq *pb.TarekomiReq) (*empty.Empty, error) {
	panic("not impl")
}

func (pth *PrivateTarekomiHandler) TarekomiBoard(c context.Context, tbordreq *pb.TarekomiBoardReq) (*pb.TarekomiBoardRes, error) {
	panic("not impl")
}

func (pth *PrivateTarekomiHandler) AddStar(c context.Context, addstarreq *pb.AddStarReq) (*empty.Empty, error) {
	panic("not impl")
}

func (pth *PrivateTarekomiHandler) GetStaredTarekomi(c context.Context, e *empty.Empty) (*pb.TarekomiSummaries, error) {
	panic("not impl")
}
