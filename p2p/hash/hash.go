package hash

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

// GenerateNonce ... ハッシュ計算のためのナンスを雑に作る
// 引数でNonceの文字列長を指定できる
func GenerateNonce(n int) string {
	b := []byte{}
	for i := 0; i < n; i++ {
		b = append(b, rs1Letters[rand.Intn(len(rs1Letters))])
	}
	return string(b)
}

// CanGenerateBlock ... ハッシュ計算をして、ブロックが作れるかどうかを見る
// この関数がTRUEを返したら、見事ブロックを作成する権利が得られる
func CanGenerateBlock(prevHash, nonce string) bool {
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
