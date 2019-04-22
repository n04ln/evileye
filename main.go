package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/NoahOrberg/evileye/controller"
	"github.com/NoahOrberg/evileye/grpcauth"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	"github.com/NoahOrberg/evileye/usecase"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
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

// impl --> public_health.go

// func (s *server) HealthCheck(c context.Context, e *empty.Empty) (*pb.HealthCheckRes, error) {
// 	hash, buildatstr := checkHealth(c)
// 	buildatunix, err := strconv.ParseUint(buildatstr, 10, 64)

// 	if err != nil {
// 		return nil, status.Error(codes.Internal, "invalid unixtime")
// 	}

// 	return &pb.HealthCheckRes{CommitHash: hash, BuildTime: buildatunix}, nil
// }

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

// func checkHealth(c context.Context) (string, string) {
// 	return commitHash, buildTime
// }

func main() {

	db, err := sqlx.Open("sqlite3", "data.sqlite3")
	if err != nil {
		panic(err)
	}

	pur := repository.NewSqliteUserRepository(db)

	puus := usecase.NewUserUsecase(pur, 100*time.Second)

	publicServerHandler := controller.NewPublicServerHandler(puus)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcauth.UnaryServerInterceptor(
				grpcauth.UserAuth(),
			),
		),
	)

	pb.RegisterPrivateServer(s, &server{})
	pb.RegisterPublicServer(s, publicServerHandler)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
