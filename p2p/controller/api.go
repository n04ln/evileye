package p2pcontroller

import (
	"context"
	"log"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type p2pServer struct {
	clis map[string]pb.InternalClient //NOTE: map[HOST]*Client
}

// NewP2PServer is a constructor for p2p service. (for internal conversation)
func NewP2PServer(hosts []string) (pb.InternalServer, error) {
	// Connect other nodes
	clis := make(map[string]pb.InternalClient)
	for _, host := range hosts {
		conn, err := grpc.Dial(host, grpc.WithInsecure())
		if err != nil {
			log.Printf("[INFO] did not connect: %v", err)
		}
		clis[host] = pb.NewInternalClient(conn)
	}

	return &p2pServer{
		clis: clis,
	}, nil
}

func (s *p2pServer) SuccessHashCalc(context.Context, *pb.SuccessHashCalcRequest) (*empty.Empty, error) {
	panic("not impl yet")
}

func (s *p2pServer) SendCheckResult(context.Context, *pb.SendCheckResultRequest) (*empty.Empty, error) {
	panic("not impl yet")
}

func (s *p2pServer) GetTxPool(context.Context, *empty.Empty) (*pb.Tarekomis, error) {
	panic("not impl yet")
}
