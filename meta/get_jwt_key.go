package meta

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	authHeader = "Authorization"
)

// JWTトークンを取得する
func GetAuthorizationKey(ctx context.Context) (string, error) {
	return fromMeta(ctx, authHeader)
}

func fromMeta(ctx context.Context, key string) (string, error) {
	vs, ok := ctx.Value(authHeader).(string)
	if !ok {
		return "", status.Error(codes.FailedPrecondition, "invalid authorization key")
	}

	parts := strings.SplitN(vs, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", fmt.Errorf("invalid auth header")
	}

	return parts[1], nil
}
