package grpcauth

import (
	"context"
	"errors"
	"time"

	"github.com/NoahOrberg/evileye/meta"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	contextKey = "userinfo"
)

func UserAuth() DefaultAuthenticateFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := meta.GetAuthorizationKey(ctx)
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "")
		}

		// TODO: JWT secret を設定する
		ui, err := GetUserFromJWT(token, "")
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		if ui.ExpiredAt < time.Now().Unix() {
			return nil, status.Error(codes.DeadlineExceeded, "timeout")
		}

		ctx = context.WithValue(ctx, contextKey, ui)

		return ctx, nil
	}
}

type UserInfo struct {
	ID        int64
	ExpiredAt int64 // unixtime
}

type UserJWT struct {
	UserInfo
	jwt.StandardClaims
}

func GetUserFromJWT(token, secret string) (UserInfo, error) {
	var u UserJWT
	_, err := jwt.ParseWithClaims(token, &u, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return UserInfo{}, errors.New("failed to get username from jwt")
	}

	return u.UserInfo, nil
}
