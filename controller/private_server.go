package controller

import (
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/usecase"
)

type privateServer struct {
	TUsecase usecase.ServerTarekomiUsecase
}

func NewPrivServer(tu usecase.ServerTarekomiUsecase) pb.PrivateServer {
	return &privateServer{TUsecase: tu}
}
