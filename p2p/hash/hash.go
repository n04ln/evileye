package p2phash

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	uuid "github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// BackgroundTask includes repository and clients information
type BackgroundTask struct {
	repo repository.Blocks // TODO: define it
	clis []pb.InternalClient
}

func NewBackgroundTask(
	hosts []string, repo repository.Blocks) (*BackgroundTask, error) {

	clis := make([]pb.InternalClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		clis = append(clis, pb.NewInternalClient(conn))
	}

	return &BackgroundTask{
		repo: repo,
		clis: clis,
	}, nil
}

// Do calculate Hash in background
// if it is success, then broadcast it using `SuccessHashCalc` to all nodes.
// NOTE: Should do `go b.Do()` in main function
func (b *BackgroundTask) Do() {
	for {
		nonce := generateNonce(rand.Intn(6)) // NOTE: 6 is default, can change it but if it is large, calclation is to slow.
		latestBlock, err := b.repo.GetLatestBlock()
		if err != nil {
			log.L().Error("failed blockRepo.GetLatestBlockHash", zap.Error(err))
			continue
		}
		prevHash := latestBlock.Hash
		if canGenerateBlock(prevHash, nonce) {
			id := uuid.New() // NOTE: save it? maybe ok that is not necessary save.
			for _, cli := range b.clis {
				_, err := cli.SuccessHashCalc(context.Background(), &pb.SuccessHashCalcRequest{
					Id:    id.String(),
					Nonce: nonce,
				})
				if err != nil {
					log.L().Error("failed SuccessHashCalc",
						zap.Error(err),
						zap.String("id", id.String()),
						zap.String("nonce", nonce),
					)
				}
			}
		}
		time.Sleep(100 * time.Millisecond) // NOTE: sloppy sleep
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
		"HEI",
		"SEI",
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
		zap.Strings("parts", parts),
		zap.Bool("isOk", isOk)) // NOTE: [HEI, SEI] が入ってることをわかりやすくしたい
	return isOk
}
