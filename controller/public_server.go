package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/usecase"
)

type publicServer struct {
	commitHash string
	buildTime  string
	UUsecase   usecase.ServerUserUsecase
}

func NewPublicServer(commitHash, buildTime string, su usecase.ServerUserUsecase) pb.PublicServer {
	return &publicServer{
		commitHash: commitHash,
		buildTime:  buildTime,
		UUsecase:   su,
	}
}
