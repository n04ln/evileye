package main

import (
	"net"
	"os"
	"time"

	"github.com/NoahOrberg/evileye/controller"
	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	p2phash "github.com/NoahOrberg/evileye/p2p/hash"
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
	ur := repository.NewSqliteUserRepository(db)

	publicServer := controller.NewPublicServer(commitHash, buildTime, puus)
	privServer := controller.NewPrivServer(tr, sr, vr, ur)

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

	driverName := "sqlite3"
	dbPath := "./data.sqlite3"
	db, err = sqlx.Open(driverName, dbPath) // TODO: maybe path is invalid in container.
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
	log.L().Info("Please wait 30 seconds for provisioning node server!")
	// NOTE: for provisioning
	time.Sleep(30 * time.Second)
	go bTask.Do() // NOTE: this task is calc hash infinity :)

	log.L().Info("Start Server And Node!")
	if err := s.Serve(lis); err != nil {
		log.L().Fatal("failed to serve", zap.Error(err))
	}
}
