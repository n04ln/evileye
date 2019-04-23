package controller

import (
	"context"
	"strconv"

	pb "github.com/NoahOrberg/evileye/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	commitHash string
	buildTime  string
)

func checkHealth(c context.Context) (string, string) {
	return commitHash, buildTime
}

func (psh *publicServer) HealthCheck(c context.Context, e *pb.Empty) (*pb.HealthCheckRes, error) {
	hash, buildatstr := checkHealth(c)
	buildatunix, err := strconv.ParseUint(buildatstr, 10, 64)

	if err != nil {
		return nil, status.Error(codes.Internal, "invalid unixtime")
	}

	return &pb.HealthCheckRes{CommitHash: hash, BuildTime: buildatunix}, nil
}
