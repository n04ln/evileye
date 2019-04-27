package p2pcontroller

import (
	"context"
	"sync"

	"github.com/NoahOrberg/evileye/log"
	"github.com/NoahOrberg/evileye/p2p/hash"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type p2pServer struct {
	txPool []struct {
		UserName string `json:"user_name"`
		URL      string `json:"url"`
	}
	b    p2phash.BackgroundTask
	repo repository.Blocks // TODO: define it
	// NOTE: clis[0] is LEADER, So Arrays order must be same between each nodes.
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
			log.L().Error("did not connect: %v", zap.Error(err))
		}
		clis[host] = pb.NewInternalClient(conn)
	}

	return &p2pServer{
		clis: clis,
	}, nil
}

func (s *p2pServer) SuccessHashCalc(ctx context.Context, req *pb.SuccessHashCalcRequest) (*empty.Empty, error) {
	nonce := req.GetNonce() // なんす
	id := req.GetId()       // リクエストのID
	var ok bool
	if ok = s.b.IsValidNonce(nonce); !ok {
		log.L().Error("invalid nonce",
			zap.String("id", id),
			zap.String("nonce", nonce))
	} else {
		log.L().Error("VALID NONCE !!!",
			zap.String("id", id),
			zap.String("nonce", nonce))
	}
	for _, cli := range s.clis {
		if _, err := cli.SendCheckResult(ctx, &pb.SendCheckResultRequest{
			Id:   id,
			IsOk: ok,
		}); err != nil {
			log.L().Error("cannot send to other node",
				zap.String("id", id),
				zap.String("nonce", nonce))
		}
	}
	return &empty.Empty{}, nil
}

// 他のノードが計算確認したものを受け取る
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
	}

	return &empty.Empty{}, nil
}

func (s *p2pServer) GetTxPool(context.Context, *empty.Empty) (*pb.Tarekomis, error) {
	// TODO: リーダーのみうけとることができる
	return nil, nil
}
