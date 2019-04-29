package controller

import (
	"bytes"
	"context"
	"crypto/sha256"

	"github.com/NoahOrberg/evileye/config"
	"github.com/NoahOrberg/evileye/jwt"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (psh *publicServer) Login(c context.Context, loginreq *pb.LoginRequest) (*pb.LoginRes, error) {
	u, err := psh.UUsecase.UserGetByName(c, loginreq.ScreenName)
	if err != nil {
		log.L().Error("UserGetByName failed", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "user name not match")
	}

	ok := rawpasswdcomp(string(u.Password), loginreq.Password)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user password not match")
	}

	// TODO: jwt token secret は後で設定する
	token, err := jwt.CreateJWTToken(*u, config.GetConfig().Secret)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "cannnot set jwt token")
	}

	return &pb.LoginRes{Token: token}, nil
}

func encryptPassword(s string) []byte {
	salt := config.GetConfig().Salt
	raw := []byte(s)

	res := sha256.Sum256(append(raw, salt...))
	return res[:]
}

func isCorrectPassword(encripted []byte, rawpw string) bool {
	enc := encryptPassword(rawpw)

	return bytes.Equal(encripted, enc)
}

func rawpasswdcomp(pw, rawpw string) bool {
	return pw == rawpw
}
