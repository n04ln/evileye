package p2pcontroller

import (
	"context"
	"encoding/json"
	"errors"
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
	leaderHost = "evileye1:50051"
)

type p2pServer struct {
	// リーダーノード用
	txPool  []*pb.Tx            // NOTE: とりあえずここにいれる
	waitTxs map[string][]*pb.Tx // NOTE: SuccessHashCalcのタイミングでこっちに入れ直す

	b    *p2phash.BackgroundTask
	repo repository.Blocks // TODO: define it

	once  *sync.Once
	hosts []string

	// NOTE: clis[0] is LEADER, So Arrays order must be same between each nodes.
	clis           map[string]pb.InternalClient // NOTE: map[HOST]*Client
	successHashCnt map[string]int64             // NOTE: MUST USE RWMutex
	failedHashCnt  map[string]int64
	mux            *sync.RWMutex
}

// NewP2PServer is a constructor for p2p service. (for internal conversation)
func NewP2PServer(hosts []string, b *p2phash.BackgroundTask, repo repository.Blocks) (pb.InternalServer, error) {
	return &p2pServer{
		hosts:          hosts,
		b:              b,
		repo:           repo,
		waitTxs:        make(map[string][]*pb.Tx),
		txPool:         make([]*pb.Tx, 0, 100),
		successHashCnt: map[string]int64{},
		failedHashCnt:  map[string]int64{},
		once:           new(sync.Once),
		mux:            new(sync.RWMutex),
	}, nil
}

func (s *p2pServer) HealthCheck(ctx context.Context, req *empty.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Hash: "hash_dayo",
	}, nil
}

func (s *p2pServer) SuccessHashCalc(ctx context.Context, req *pb.SuccessHashCalcRequest) (*empty.Empty, error) {
	log.L().Info("Invoked SuccessHashCalc", zap.Any("req", req))
	log.L().Info("SEND BACKGROUND TASK SLEEP SIGNAL")
	p2phash.StopCalc <- struct{}{}

	log.L().Info("Do s.once.Do")
	s.once.Do(func() {
		log.L().Info("FIRST SUCCESS HASH CALC, So add Client Connection!")
		// Connect other nodes
		clis := make(map[string]pb.InternalClient)
		for _, host := range s.hosts {
			conn, err := grpc.Dial(host, grpc.WithInsecure())
			if err != nil {
				log.L().Error("did not connect: %v", zap.Error(err))
			}
			clis[host] = pb.NewInternalClient(conn)
		}
		s.clis = clis
	})
	log.L().Info("Done s.once.Do")

	// (やるのはリーダーノードだけでいい)
	s.mutexOpe(func() {
		s.waitTxs[req.GetId()] = s.txPool
		s.txPool = make([]*pb.Tx, 0, 100)
		log.L().Info("saved txpool contents to waitTxs",
			zap.Any("waitTxs", s.waitTxs),
			zap.Any("txPool", s.txPool))
	})

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

	s.mutexOpe(func() {
		waitCh := make(chan struct{}, len(s.clis))
		done := make(chan struct{}, 1)
		for host, cli := range s.clis {
			go func(host string, cli pb.InternalClient) {
				waitCh <- struct{}{}
				<-done
				log.L().Info("Send SendCheckResult",
					zap.String("host", host),
					zap.String("id", id),
					zap.String("nonce", nonce))
				if _, err := cli.SendCheckResult(
					context.Background(), &pb.SendCheckResultRequest{
						Id:    id,
						Nonce: nonce,
						IsOk:  ok,
					}); err != nil {
					log.L().Error("cannot send to other node",
						zap.String("id", id),
						zap.String("nonce", nonce),
						zap.Error(err))
				}
				log.L().Info("SENDED SendCheckResult")
			}(host, cli)
		}
		for i := 0; i < len(s.clis); i++ {
			<-waitCh
			log.L().Info("waiting for Sending SuccessHashCalc...")
		}
		close(done) // DONE!!
	})

	return &empty.Empty{}, nil
}

func (s *p2pServer) mutexOpe(f func()) {
	log.L().Info("Do s.Lock")
	s.mux.Lock()
	log.L().Info("Done s.Lock")
	f()
	log.L().Info("Do s.Unlock")
	s.mux.Unlock()
	log.L().Info("Done s.Unlock")
}

func (s *p2pServer) mutexOpeWithErr(f func() error) error {
	log.L().Info("Do s.Lock")
	s.mux.Lock()
	log.L().Info("Done s.Lock")
	err := f()
	log.L().Info("Do s.Unlock")
	s.mux.Unlock()
	log.L().Info("Done s.Unlock")
	return err
}

// 他のノードが計算確認したものを受け取る
func (s *p2pServer) SendCheckResult(ctx context.Context, req *pb.SendCheckResultRequest) (*empty.Empty, error) {
	log.L().Info("Invoked SendCheckResult", zap.Any("req", req))

	// NOTE: なければ初期化
	s.mutexOpe(func() {
		if _, ok := s.successHashCnt[req.GetId()]; !ok {
			s.successHashCnt[req.GetId()] = 0
		}
	})
	s.mutexOpe(func() {
		if _, ok := s.failedHashCnt[req.GetId()]; !ok {
			s.failedHashCnt[req.GetId()] = 0
		}
	})

	if req.GetIsOk() {
		s.mutexOpe(func() {
			s.successHashCnt[req.GetId()]++
		})
	} else {
		s.mutexOpe(func() {
			s.failedHashCnt[req.GetId()]++
		})
	}

	s.mutexOpe(func() {
		log.L().Info("NOW, VOTING INFO",
			zap.Int64("successCnt", s.successHashCnt[req.GetId()]),
			zap.Int64("failedCnt", s.failedHashCnt[req.GetId()]))
	})

	// 承認数が明らかに無理になったらTxPoolに戻す

	if err := s.mutexOpeWithErr(func() error {
		if s.failedHashCnt[req.GetId()] >= 2 {
			// NOTE: 今回は順番などは関係ないので、Orderingを意識して戻すのは不要
			p2phash.RestartCalc <- struct{}{}

			s.txPool = append(s.txPool, s.waitTxs[req.GetId()]...)
			delete(s.waitTxs, req.GetId())
			return errors.New("break")
		}
		return nil
	}); err != nil {
		return &empty.Empty{}, nil
	}

	if err := s.mutexOpeWithErr(func() error {
		if s.successHashCnt[req.GetId()] >= 2 /* しきい値を環境変数注入 */ {
			s.successHashCnt[req.GetId()] = -999 // NOTE: 雑に複数回Insertされないようにする
			p2phash.RestartCalc <- struct{}{}

			s.once.Do(func() {
				log.L().Info("FIRST SUCCESS HASH CALC, So add Client Connection!")
				// Connect other nodes
				clis := make(map[string]pb.InternalClient)
				for _, host := range s.hosts {
					conn, err := grpc.Dial(host, grpc.WithInsecure())
					if err != nil {
						log.L().Error("did not connect: %v", zap.Error(err))
					}
					clis[host] = pb.NewInternalClient(conn)
				}
				s.clis = clis
			})

			log.L().Info("GetTxPool from leaderHost")
			txs, err := s.clis[leaderHost].GetTxPool(
				context.Background(),
				&pb.GetTxPoolRequest{Id: req.GetId()})
			if err != nil {
				log.L().Error("cannot GetTxPool",
					zap.Error(err))
				return err
			}
			data, err := json.Marshal(txs)
			if err != nil {
				log.L().Error("cannot Marshal Txs",
					zap.Error(err))
				return err
			}
			prevBlock, err := s.repo.GetLatestBlock()
			if err != nil {
				log.L().Error("cannot Marshal Txs",
					zap.Error(err))
				return err
			}
			log.L().Info("Add BLOCK!!!")
			log.L().Info("Calc Hash for InsertBlock")
			hash := p2phash.CalcHash(prevBlock.Hash, req.GetNonce())
			log.L().Info("insertBlock",
				zap.String("prevHash", prevBlock.Hash),
				zap.String("data", string(data)),
				zap.String("hash", hash))
			block, err := s.repo.InsertBlock(ctx, string(data), prevBlock.Hash, hash)
			if err != nil {
				log.L().Error("cannot InsertBlock",
					zap.Error(err))
				return err
			}
			if block != nil {
				log.L().Info("Successfully Add Block",
					zap.Int64("id", block.ID),
					zap.Int64("create_time", block.CreateTime),
					zap.String("prevHash", block.PrevHash),
					zap.String("data", block.Data),
					zap.String("hash", block.Hash))
			}
			log.L().Info("finish add block")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *p2pServer) GetTxPool(ctx context.Context, req *pb.GetTxPoolRequest) (*pb.Txs, error) {
	log.L().Info("Invoked GetTxPool", zap.Any("req", req))
	txs := make([]*pb.Tx, 0, len(s.waitTxs[req.GetId()]))
	for _, tx := range s.waitTxs[req.GetId()] {
		txs = append(txs, &pb.Tx{
			UserName:          tx.GetUserName(),
			Url:               tx.GetUrl(),
			Desc:              tx.GetDesc(),
			ApprovedUserNames: tx.GetApprovedUserNames(),
		})
	}

	log.L().Info("finish GetTxPool")
	return &pb.Txs{
		Txs: txs,
	}, nil
}

func (s *p2pServer) SendTx(ctx context.Context, req *pb.SendTxRequest) (*empty.Empty, error) {
	log.L().Info("Invoked SendTx", zap.Any("req", req))
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
