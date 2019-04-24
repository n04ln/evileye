package p2pcontroller

import (
	"context"
	"log"
	"sync"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type p2pServer struct {
	clis            map[string]pb.InternalClient // NOTE: map[HOST]*Client
	successHashCalc map[string]int64             // NOTE: MUST USE RWMutex
	mux             sync.RWMutex
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

func (s *p2pServer) SuccessHashCalc(ctx context.Context, req *pb.SuccessHashCalcRequest) (*empty.Empty, error) {
	// 再計算してできたらSendCheckResultにブロードキャスト
	panic("not impl")
}

func (s *p2pServer) SendCheckResult(ctx context.Context, req *pb.SendCheckResultRequest) (*empty.Empty, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	// NOTE: なければ初期化
	if _, ok := s.successHashCalc[req.GetId()]; !ok {
		s.successHashCalc[req.GetId()] = 0
	}

	if req.GetIsOk() {
		s.successHashCalc[req.GetId()]++
	}

	if s.successHashCalc[req.GetId()] >= 3 /* しきい値を環境変数注入 */ {
		// TODO: Block追加
	}

	return &empty.Empty{}, nil
}

func (s *p2pServer) GetTxPool(context.Context, *empty.Empty) (*pb.Tarekomis, error) {
	// TODO: リーダーのみうけとることができる

}
