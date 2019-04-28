package p2phash

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"sync"
	"time"

	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	uuid "github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	RestartCalc = make(chan struct{}, 5)
	StopCalc    = make(chan struct{}, 5)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// BackgroundTask includes repository and clients information
type BackgroundTask struct {
	repo  repository.Blocks // TODO: define it
	clis  map[string]pb.InternalClient
	once  *sync.Once
	isDo  bool
	hosts []string
}

func NewBackgroundTask(
	hosts []string, repo repository.Blocks) (*BackgroundTask, error) {

	return &BackgroundTask{
		repo:  repo,
		hosts: hosts,
		once:  new(sync.Once),
	}, nil
}

// Do calculate Hash in background
// if it is success, then broadcast it using `SuccessHashCalc` to all nodes.
// NOTE: Should do `go b.Do()` in main function
func (b *BackgroundTask) Do() {
	for {
		select {
		case <-StopCalc:
			b.isDo = false
			log.L().Info("BACKGROUND TASK SLEEP!")
			go func() {
				time.Sleep(5 * time.Second)
				if !b.isDo {
					log.L().Info("FORCE RESTART!")
					RestartCalc <- struct{}{}
					return
				}
				log.L().Info("Unncessary FORCE RESTART!")
			}()
			<-RestartCalc
			b.isDo = true
			log.L().Info("BACKGROUND TASK RESTART!")
		default:
		}
		nonce := generateNonce(rand.Intn(6)) // NOTE: 6 is default, can change it but if it is large, calclation is to slow.
		latestBlock, err := b.repo.GetLatestBlock()
		if err != nil {
			log.L().Error("failed blockRepo.GetLatestBlockHash", zap.Error(err))
			continue
		}
		prevHash := latestBlock.Hash
		if canGenerateBlock(prevHash, nonce) {
			b.once.Do(func() {
				log.L().Info("FIRST SUCESS HASH CALC in BACKGROUND TASK, So add Client Connection!")
				// Connect other nodes
				clis := make(map[string]pb.InternalClient)
				for _, host := range b.hosts {
					conn, err := grpc.Dial(host, grpc.WithInsecure())
					if err != nil {
						log.L().Error("did not connect: %v", zap.Error(err))
					}
					clis[host] = pb.NewInternalClient(conn)
				}
				b.clis = clis
			})

			id := uuid.New() // NOTE: save it? maybe ok that is not necessary save.
			waitCh := make(chan struct{}, len(b.clis))
			done := make(chan struct{}, 1)
			for host, cli := range b.clis {
				go func(host string, cli pb.InternalClient) {
					waitCh <- struct{}{}
					<-done
					log.L().Info("SEND SuccessHashCalc",
						zap.String("host", host),
						zap.String("id", id.String()),
						zap.String("nonce", nonce))
					_, err := cli.SuccessHashCalc(
						context.Background(), &pb.SuccessHashCalcRequest{
							Id:    id.String(),
							Nonce: nonce,
						})
					log.L().Info("SENDED SuccessHashCalc")
					if err != nil {
						log.L().Error("failed SuccessHashCalc",
							zap.Error(err),
							zap.String("id", id.String()),
							zap.String("nonce", nonce),
						)
					}
				}(host, cli)
			}
			for i := 0; i < len(b.clis); i++ {
				<-waitCh
				log.L().Info("waiting for Sending SuccessHashCalc...")
			}
			close(done) // DONE!!
		}
		time.Sleep(10 * time.Millisecond) // NOTE: sloppy sleep
	}
}

// generateNonce ... ハッシュ計算のためのナンスを雑に作る
// 引数でNonceの文字列長を指定できる
func generateNonce(n int) string {
	b := []byte{}
	for i := 0; i < n; i++ {
		b = append(b, rs1Letters[rand.Intn(len(rs1Letters))])
	}
	return string(b)
}

// IsValidNonce is
func (b BackgroundTask) IsValidNonce(nonce string) bool {
	latestBlock, err := b.repo.GetLatestBlock()
	if err != nil {
		log.L().Error("failed blockRepo.GetLatestBlockHash", zap.Error(err))
		return false
	}
	prevHash := latestBlock.Hash
	return canGenerateBlock(prevHash, nonce)
}

// canGenerateBlock ... ハッシュ計算をして、ブロックが作れるかどうかを見る
// この関数がTRUEを返したら、見事ブロックを作成する権利が得られる
func canGenerateBlock(prevHash, nonce string) bool {
	// このなかのいづれかが含まれればよい
	parts := []string{
		"HE",
		"II",
		"SE",
		"II",
	}

	// チェック補助関数
	check := func(hash []byte, heiseis []string) bool {
		for _, part := range parts {
			if bytes.Count(hash, []byte(part)) >= 1 {
				return true
			}
		}
		return false
	}

	// SHA256にかける
	h := sha256.Sum256([]byte(prevHash + nonce))
	isOk := check(h[:], parts)
	log.L().Info("calced Hash is",
		zap.String("value", hex.EncodeToString(h[:])),
		zap.String("nonce", nonce),
		zap.Strings("parts", parts),
		zap.Bool("isOk", isOk)) // NOTE: [HEI, SEI] が入ってることをわかりやすくしたい
	return isOk
}

func CalcHash(prevHash, nonce string) string {
	h := sha256.Sum256([]byte(prevHash + nonce))
	return hex.EncodeToString(h[:])
}
