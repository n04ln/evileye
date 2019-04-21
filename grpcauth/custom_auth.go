package grpcauth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ref : https://qiita.com/yoheimuta/items/72d4b75f72d8913adc10
type DefaultAuthenticateFunc func(ctx context.Context) (context.Context, error)

func UnaryServerInterceptor(authFunc DefaultAuthenticateFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := authFunc(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return handler(newCtx, req)
	}
}
