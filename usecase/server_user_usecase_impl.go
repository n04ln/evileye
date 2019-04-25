package usecase

import (
	"context"
	"time"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/repository"
)

type ServerUserUsecaseImpl struct {
	userRepo       repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(ur repository.UserRepository, ct time.Duration) ServerUserUsecase {
	return &ServerUserUsecaseImpl{
		userRepo:       ur,
		contextTimeout: ct,
	}
}

func (u *ServerUserUsecaseImpl) UserGetByID(c context.Context, id int64) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.UserGetByID(ctx, id)
}

func (u *ServerUserUsecaseImpl) UserGetByIDList(c context.Context, limit, offset int64) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.UserGetByIDList(ctx, limit, offset)
}

func (u *ServerUserUsecaseImpl) UserGetByName(c context.Context, uname string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.UserGetByName(ctx, uname)
}

func (u *ServerUserUsecaseImpl) Store(c context.Context, usr *entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.Store(ctx, usr)
}
