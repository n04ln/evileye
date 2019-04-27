package usecase

import (
	"context"
	"time"

	"github.com/NoahOrberg/evileye/entity"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/NoahOrberg/evileye/repository"
)

type ServerTarekomiUsecaseImpl struct {
	tarekomiRepo   repository.TarekomiRepository
	contextTimeout time.Duration
}

func NewTarekomiUsecase(tr repository.TarekomiRepository, ct time.Duration) ServerTarekomiUsecase {
	return &ServerTarekomiUsecaseImpl{
		tarekomiRepo:   tr,
		contextTimeout: ct,
	}
}

func (t *ServerTarekomiUsecaseImpl) GetTarekomiFromUser(c context.Context, id int64, limit int64, offset int64) (pb.TarekomiSummaries, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	return t.tarekomiRepo.GetTarekomiFromUser(ctx, id, limit, offset)
}
func (t *ServerTarekomiUsecaseImpl) GetTarekomiBoard(c context.Context, limit int64, offset int64) (pb.TarekomiSummaries, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	return t.tarekomiRepo.GetTarekomiBoard(ctx, limit, offset)
}
func (t *ServerTarekomiUsecaseImpl) Store(c context.Context, tk entity.Tarekomi) (int64, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	return t.tarekomiRepo.Store(ctx, tk)
}
func (t *ServerTarekomiUsecaseImpl) UpdateTarekomiState(c context.Context, nt entity.Tarekomi) (entity.Tarekomi, error) {
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	return t.tarekomiRepo.UpdateTarekomiState(ctx, nt)
}
