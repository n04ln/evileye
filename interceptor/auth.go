package interceptor

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/NoahOrberg/evileye/jwt"
	"github.com/NoahOrberg/evileye/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	authHeader = "authorization"
	contextKey = new(int)
)

func WithJWT(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	method, _ := grpc.Method(ctx)
	spMethod := strings.Split(method, "/")
	log.L().Info("show method", zap.String("method", method))
	if len(spMethod) < 2 {
		log.L().Error("invalid method", zap.String("method", method))
		return nil, status.Error(
			codes.InvalidArgument, fmt.Sprintf("method is invalid: %s", method))
	}

	if spMethod[1] == "evileye.Public" {
		return handler(ctx, req)
	}

	log.L().Info("necessary Authorization")
	token, err := getAuthorizationKey(ctx)
	if err != nil {
		log.L().Error("meta.GetAuthorizationKey is failed", zap.Error(err))
		return nil, status.Error(codes.FailedPrecondition, "cannnot get token from metadata")
	}
	log.L().Info("show token", zap.String("token", token))

	// TODO: JWT secret を設定する
	ui, err := jwt.GetUserInfoFromJWT(token, "")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if ui.ExpiredAt < time.Now().Unix() {
		return nil, status.Error(codes.Unauthenticated, "timeout")
	}

	ctx = context.WithValue(ctx, contextKey, ui)
	return handler(ctx, req)
}

// JWTトークンを取得する
func getAuthorizationKey(ctx context.Context) (string, error) {
	return fromMeta(ctx, authHeader)
}

func fromMeta(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("cannot get metadata")
	}
	vs := md[key]
	if len(vs) == 0 {
		return "", errors.New("token is not set")
	}

	parts := strings.SplitN(vs[0], " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", fmt.Errorf("invalid auth header")
	}

	return parts[1], nil
}
