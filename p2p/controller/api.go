package p2pcontroller

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/NoahOrberg/evileye/log"
	"github.com/NoahOrberg/evileye/p2p/hash"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	leaderHost = "evileye:50051"
)

type p2pServer struct {
	// リーダーノード用
	txPool  []*pb.Tx // NOTE: とりあえずここにいれる
	waitTxs []*pb.Tx // NOTE: SuccessHashCalcのタイミングでこっちに入れ直す

	b    *p2phash.BackgroundTask
	repo repository.Blocks // TODO: define it

	// NOTE: clis[0] is LEADER, So Arrays order must be same between each nodes.
	clis           map[string]pb.InternalClient // NOTE: map[HOST]*Client
	successHashCnt map[string]int64             // NOTE: MUST USE RWMutex
	failedHashCnt  map[string]int64
	mux            sync.RWMutex
}

// NewP2PServer is a constructor for p2p service. (for internal conversation)
func NewP2PServer(hosts []string, b *p2phash.BackgroundTask, repo repository.Blocks) (pb.InternalServer, error) {
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
		clis:           clis,
		b:              b,
		repo:           repo,
		waitTxs:        make([]*pb.Tx, 0, 100),
		txPool:         make([]*pb.Tx, 0, 100),
		successHashCnt: map[string]int64{},
		failedHashCnt:  map[string]int64{},
	}, nil
}

func (s *p2pServer) SuccessHashCalc(ctx context.Context, req *pb.SuccessHashCalcRequest) (*empty.Empty, error) {
	log.L().Info("Invoved SuccessHashCalc", zap.Any("req", req))

	s.mux.Lock()
	defer s.mux.Unlock()

	s.waitTxs = s.txPool
	s.txPool = make([]*pb.Tx, 0, 100)

	nonce := req.GetNonce() // なんす
	id := req.GetId()       // リクエストのID
	var ok bool
	if ok = s.b.IsValidNonce(nonce); !ok {
		log.L().Error("INVALID nonce",
			zap.String("id", id),
			zap.String("nonce", nonce))
	} else {
		log.L().Info("VALID NONCE !!!",
			zap.String("id", id),
			zap.String("nonce", nonce))
	}
	for _, cli := range s.clis {
		if _, err := cli.SendCheckResult(ctx, &pb.SendCheckResultRequest{
			Id:    id,
			Nonce: nonce,
			IsOk:  ok,
		}); err != nil {
			log.L().Error("cannot send to other node",
				zap.String("id", id),
				zap.String("nonce", nonce),
				zap.Error(err))
		}
	}
	return &empty.Empty{}, nil
}

// 他のノードが計算確認したものを受け取る
func (s *p2pServer) SendCheckResult(ctx context.Context, req *pb.SendCheckResultRequest) (*empty.Empty, error) {
	log.L().Info("Invoved SendCheckResult", zap.Any("req", req))
	s.mux.Lock()
	defer s.mux.Unlock()

	// NOTE: なければ初期化
	if _, ok := s.successHashCnt[req.GetId()]; !ok {
		s.successHashCnt[req.GetId()] = 0
	}
	if _, ok := s.failedHashCnt[req.GetId()]; !ok {
		s.failedHashCnt[req.GetId()] = 0
	}

	if req.GetIsOk() {
		s.successHashCnt[req.GetId()]++
	} else {
		s.failedHashCnt[req.GetId()]++
	}

	// 承認数が明らかに無理になったらTxPoolに戻す
	if s.failedHashCnt[req.GetId()] >= 2 {
		// NOTE: 今回は順番などは関係ないので、Orderingを意識して戻すのは不要
		s.txPool = append(s.txPool, s.waitTxs...)
	}

	if s.successHashCnt[req.GetId()] >= 2 /* しきい値を環境変数注入 */ {
		txs, err := s.clis[leaderHost].GetTxPool(ctx, &empty.Empty{})
		if err != nil {
			log.L().Error("cannot GetTxPool",
				zap.Error(err))
			return nil, err
		}
		data, err := json.Marshal(txs)
		if err != nil {
			log.L().Error("cannot Marshal Txs",
				zap.Error(err))
			return nil, err
		}
		prevBlock, err := s.repo.GetLatestBlock()
		if err != nil {
			log.L().Error("cannot Marshal Txs",
				zap.Error(err))
			return nil, err
		}
		hash := p2phash.CalcHash(prevBlock.Hash, req.GetNonce())
		_, err = s.repo.InsertBlock(ctx, string(data), prevBlock.Hash, hash)
		if err != nil {
			log.L().Error("cannot InsertBlock",
				zap.Error(err))
		}
	}

	return &empty.Empty{}, nil
}

func (s *p2pServer) GetTxPool(ctx context.Context, req *empty.Empty) (*pb.Txs, error) {
	log.L().Info("Invoved GetTxPool", zap.Any("req", req))
	txs := make([]*pb.Tx, 0, len(s.waitTxs))
	for _, tx := range s.waitTxs {
		txs = append(txs, &pb.Tx{
			UserName:          tx.GetUserName(),
			Url:               tx.GetUrl(),
			Desc:              tx.GetDesc(),
			ApprovedUserNames: tx.GetApprovedUserNames(),
		})
	}

	return &pb.Txs{
		Txs: txs,
	}, nil
}

func (s *p2pServer) SendTx(ctx context.Context, req *pb.SendTxRequest) (*empty.Empty, error) {
	log.L().Info("Invoved SendTx", zap.Any("req", req))
	s.mux.Lock()
	defer s.mux.Unlock()

	s.txPool = append(s.txPool, &pb.Tx{
		UserName:          req.GetUserName(),
		Url:               req.GetUrl(),
		Desc:              req.GetDesc(),
		ApprovedUserNames: req.GetApprovedUserNames(),
	})

	return &empty.Empty{}, nil
}
