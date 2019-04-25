package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
)

type privateServer struct {
	TRepository repository.TarekomiRepository
	SRepository repository.StarRepository
}

func NewPrivServer(tr repository.TarekomiRepository, sr repository.StarRepository) pb.PrivateServer {
	return &privateServer{
		TRepository: tr,
		SRepository: sr,
	}
}
