package p2pclient

import (
	"context"

	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	leaderHost = "localhost:50051"
)

func NewInternalClient() (InternalClient, error) {
	conn, err := grpc.Dial(leaderHost, grpc.WithInsecure())
	if err != nil {
		log.L().Error("did not connect: %v", zap.Error(err))
	}

	return InternalClient{cli: pb.NewInternalClient(conn)}, nil
}

type InternalClient struct {
	cli pb.InternalClient // NOTE: map[HOST]*Client
}

func (i *InternalClient) SentTxToLeaderNode(
	ctx context.Context, username, url, desc string, approvedUserNames []string) error {
	if _, err := i.cli.SendTx(ctx, &pb.SendTxRequest{
		UserName:          username,
		Url:               url,
		Desc:              desc,
		ApprovedUserNames: approvedUserNames,
	}); err != nil {
		log.L().Error("cannot send Tx to leader node",
			zap.String("username", username),
			zap.String("url", url),
			zap.String("desc", desc),
			zap.Strings("approvedUserName", approvedUserNames))
		return err
	}

	return nil
}
