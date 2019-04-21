package p2phash

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// BackgroundTask includes repository and clients information
type BackgroundTask struct {
	// repo *repository.Block // TODO: define it
	// clis // TODO: define it
}

// Do calculate Hash in background
// if it is success, then broadcast it using `SuccessHashCalc` to all nodes.
// NOTE: Should do `go b.Do()` in main function
func (b *BackgroundTask) Do() {
	for {
		nonce := generateNonce(rand.Intn(6))   // NOTE: 6 is default, can change it but if it is large, calclation is to slow.
		prevHash := ""                         // TODO: get prevHash using repository in receiver
		if canGenerateBlock(prevHash, nonce) { // TODO: replace empty string to prevHash
			// TODO: Broadcast using SuccessHashCalc to all nodes. and generate UUID for requestID here.
		}
		time.Sleep(1 * time.Millisecond) // NOTE: sloppy sleep
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
	return check(h[:], parts)
}
