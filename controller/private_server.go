package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
)

type privateServer struct {
	TRepository repository.TarekomiRepository
	SRepository repository.StarRepository
	VRepository repository.SqliteVoteRepository
}

func NewPrivServer(tr repository.TarekomiRepository, sr repository.StarRepository, vr repository.SqliteVoteRepository) pb.PrivateServer {
	return &privateServer{
		TRepository: tr,
		SRepository: sr,
		VRepository: vr,
	}
}
