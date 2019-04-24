package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/usecase"
)

type publicServer struct {
	UUsecase usecase.ServerUserUsecase
}

func NewPublicServer(su usecase.ServerUserUsecase) pb.PublicServer {
	return &publicServer{
		UUsecase: su,
	}
}
