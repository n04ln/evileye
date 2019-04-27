package controller

import (
	p2pclient "github.com/NoahOrberg/evileye/p2p/client"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
)

type privateServer struct {
	TRepository repository.TarekomiRepository
	SRepository repository.StarRepository
	VRepository repository.SqliteVoteRepository
	URepository repository.SqliteUserRepository
	IClient     p2pclient.InternalClient
}

func NewPrivServer(tr repository.TarekomiRepository, sr repository.StarRepository, vr repository.SqliteVoteRepository, ur repository.SqliteUserRepository, ic p2pclient.InternalClient) pb.PrivateServer {
	return &privateServer{
		TRepository: tr,
		SRepository: sr,
		VRepository: vr,
		URepository: ur,
		IClient:     ic,
	}
}
