package main

import (
	"context"
	"net"
	"strconv"

	"github.com/NoahOrberg/evileye/infra/repository"
	"github.com/NoahOrberg/evileye/log"
	p2phash "github.com/NoahOrberg/evileye/p2p/hash"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

var (
	commitHash string
	buildTime  string
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) Tarekomi(ctx context.Context, in *pb.TarekomiReq) (*empty.Empty, error) {
	panic("not impl")
}

func (s *server) Vote(ctx context.Context, in *pb.VoteReq) (*empty.Empty, error) {
	panic("not impl")
}

func (s *server) TarekomiBoard(ctx context.Context, in *pb.TarekomiBoardReq) (*pb.TarekomiBoardRes, error) {
	panic("not impl")
}

func (s *server) GetUserInfo(ctx context.Context, in *pb.UserInfoReq) (*pb.User, error) {
	panic("not impl")
}

func (s *server) GetUserList(ctx context.Context, in *pb.GetUserListReq) (*pb.Users, error) {
	panic("not impl")
}

func (s *server) AddStar(ctx context.Context, in *pb.AddStarReq) (*empty.Empty, error) {
	panic("not impl")
}

func (s *server) GetStaredTarekomi(ctx context.Context, in *empty.Empty) (*pb.TarekomiSummaries, error) {
	panic("not impl")
}

type pubServer struct{}

func (p *pubServer) HealthCheck(c context.Context, e *empty.Empty) (*pb.HealthCheckRes, error) {
	hash, buildatstr := CheckHealth(c)
	buildatunix, err := strconv.ParseUint(buildatstr, 10, 64)

	if err != nil {
		return nil, status.Error(codes.Internal, "invalid unixtime")
	}

	return &pb.HealthCheckRes{CommitHash: hash, BuildTime: buildatunix}, nil
}

func (p *pubServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginRes, error) {
	panic("not impl yet")
}

func CheckHealth(c context.Context) (string, string) {
	return commitHash, buildTime
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.L().Fatal("failed to listen", zap.Error(err))
	}
	log.L().Info(
		"success net.Listen()",
		zap.String("protocol", "tcp"),
		zap.String("port", port))

	s := grpc.NewServer()
	pb.RegisterPublicServer(s, &pubServer{})
	pb.RegisterPrivateServer(s, &server{})
	log.L().Info(
		"register server, serve it!")

	driverName := "sqlite3"
	dbPath := "./data.sqlite3"
	db, err := sqlx.Open(driverName, dbPath) // TODO: maybe path is invalid in container.
	if err != nil {
		log.L().Fatal("cannot open DB",
			zap.Error(err),
			zap.String("driverName", driverName),
			zap.String("dbPath", dbPath),
		)
	}
	blockRepo := repository.NewBlocksRepository(db)
	bTask, err := p2phash.NewBackgroundTask(nil, blockRepo) // TODO: fill it
	if err != nil {
		log.L().Fatal("cannot create BackgroundTask")
	}
	go bTask.Do() // NOTE: this task is calc hash infinity :)

	if err := s.Serve(lis); err != nil {
		log.L().Fatal("failed to serve", zap.Error(err))
	}
}
