package usecase

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	pb "github.com/NoahOrberg/evileye/protobuf"
)

type ServerTarekomiUsecase interface {
	GetTarekomiFromUser(context.Context, int64, int64, int64) (pb.TarekomiSummaries, error)
	GetTarekomiBoard(context.Context, int64, int64) (pb.TarekomiSummaries, error)
	Store(context.Context, entity.Tarekomi) (int64, error)
	UpdateTarekomiState(context.Context, entity.Tarekomi) (entity.Tarekomi, error)
}
