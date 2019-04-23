package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

func (pth *privateServer) Tarekomi(c context.Context, tarekomireq *pb.TarekomiReq) (*pb.Empty, error) {
	panic("not impl")
}

func (pth *privateServer) TarekomiBoard(c context.Context, tbordreq *pb.TarekomiBoardReq) (*pb.TarekomiBoardRes, error) {
	panic("not impl")
}

func (pth *privateServer) AddStar(c context.Context, addstarreq *pb.AddStarReq) (*pb.Empty, error) {
	panic("not impl")
}

func (pth *privateServer) GetStaredTarekomi(c context.Context, e *pb.Empty) (*pb.TarekomiSummaries, error) {
	panic("not impl")
}
