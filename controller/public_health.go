package controller

import (
	"context"
	"strconv"

	pb "github.com/NoahOrberg/evileye/protobuf"
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

func NewPublicCheckHealthHandler() *PublicCheckHealthHandler {
	return &PublicCheckHealthHandler{}
}

func checkHealth(c context.Context) (string, string) {
	return commitHash, buildTime
}

func (pch *PublicCheckHealthHandler) HealthCheck(c context.Context, e *empty.Empty) (*pb.HealthCheckRes, error) {
	hash, buildatstr := checkHealth(c)
	buildatunix, err := strconv.ParseUint(buildatstr, 10, 64)

	if err != nil {
		return nil, status.Error(codes.Internal, "invalid unixtime")
	}

	return &pb.HealthCheckRes{CommitHash: hash, BuildTime: buildatunix}, nil
}
