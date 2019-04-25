package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
)

func (pvh *privateServer) Vote(c context.Context, votereq *pb.VoteReq) (*pb.Empty, error) {
	panic("not impl")
}
