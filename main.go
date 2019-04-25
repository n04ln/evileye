package main

import (
	"net"
	"os"
	"time"

	"github.com/NoahOrberg/evileye/controller"
	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
	"github.com/NoahOrberg/evileye/usecase"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
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

func main() {

	db, err := sqlx.Open("sqlite3", os.Getenv("DB_FILE"))
	if err != nil {
		panic(err)
	}

	pur := repository.NewSqliteUserRepository(db)

	puus := usecase.NewUserUsecase(pur, 100*time.Second)

	tr := repository.NewSqliteTarekomiRepository(db)
	sr := repository.NewSqliteStarRepository(db)
	vr := repository.NewSqliteVoteRepository(db)

	publicServer := controller.NewPublicServer(commitHash, buildTime, puus)
	privServer := controller.NewPrivServer(tr, sr, vr)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.L().Fatal("failed to listen", zap.Error(err))
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptor.WithJWT,
		),
	)

	log.L().Info(
		"success net.Listen()",
		zap.String("protocol", "tcp"),
		zap.String("port", port))

	pb.RegisterPublicServer(s, publicServer)
	pb.RegisterPrivateServer(s, privServer)
	log.L().Info(
		"register server, serve it!")

	if err := s.Serve(lis); err != nil {
		log.L().Fatal("failed to serve", zap.Error(err))
	}
}
