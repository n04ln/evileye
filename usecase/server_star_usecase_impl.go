package usecase

import (
	"context"
	"time"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/repository"
)

type ServerStarUsecaseImpl struct {
	starRepo       repository.StarRepository
	contextTimeout time.Duration
}

func NewStarUsecase(sr repository.StarRepository, ct time.Duration) ServerStarUsecase {
	return &ServerStarUsecaseImpl{
		starRepo:       sr,
		contextTimeout: ct,
	}
}

func (s *ServerStarUsecaseImpl) GetStaredTarekomiID(c context.Context, uid int64) ([]int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	return s.starRepo.GetStaredTarekomiID(ctx, uid)
}

func (s *ServerStarUsecaseImpl) DeleteStar(c context.Context, st *entity.Star) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	return s.starRepo.DeleteStar(ctx, st)
}

func (s *ServerStarUsecaseImpl) Store(c context.Context, st *entity.Star) (int64, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	return s.starRepo.Store(ctx, st)
}
