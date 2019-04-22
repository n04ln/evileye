package controller

import (
	"bytes"
	"context"
	"crypto/sha256"

	"github.com/NoahOrberg/evileye/middleware"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PublicAccountStateHandler struct {
	UUsecase usecase.ServerUserUsecase
}

func NewPublicUserHandler(su usecase.ServerUserUsecase) *PublicAccountStateHandler {
	return &PublicAccountStateHandler{
		UUsecase: su,
	}
}

func (pas *PublicAccountStateHandler) Login(c context.Context, loginreq *pb.LoginRequest) (*pb.LoginRes, error) {
	u, err := pas.UUsecase.UserGetByName(c, loginreq.ScreenName)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "user name or password not match")
	}

	ok := isCorrectPassword(u.Password, loginreq.Password)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user name or password not match")
	}

	// TODO: jwt token secret は後で設定する
	token, err := middleware.CreateJWTToken(*u, "")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "cannnot set jwt token")
	}

	return &pb.LoginRes{Token: token}, nil
}

func encryptPassword(s string) []byte {
	salt := "" // TODO: あとで設定する
	raw := []byte(s)

	res := sha256.Sum256(append(raw, salt...))
	return res[:]
}

func isCorrectPassword(encripted []byte, rawpw string) bool {
	enc := encryptPassword(rawpw)

	return bytes.Equal(encripted, enc)
}
