package controller

import (
	"context"
	"strconv"

	evpb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	commitHash string
	buildTime  string
)

type PublicCheckHealthHandler struct {
}

func checkHealth(c context.Context) (string, string) {
	return commitHash, buildTime
}

func (pch *PublicCheckHealthHandler) HealthCheck(c context.Context, e *empty.Empty) (*evpb.HealthCheckRes, error) {
	hash, buildatstr := checkHealth(c)
	buildatunix, err := strconv.ParseUint(buildatstr, 10, 64)

	if err != nil {
		return nil, status.Error(codes.Internal, "invalid unixtime")
	}

	return &evpb.HealthCheckRes{CommitHash: hash, BuildTime: buildatunix}, nil
}
