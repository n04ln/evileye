package controller

import (
	"context"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

type PrivateVoteHandler struct {
}

func (pvh *PrivateVoteHandler) Vote(c context.Context, votereq *pb.VoteReq) (*empty.Empty, error) {
	panic("not impl")
}
