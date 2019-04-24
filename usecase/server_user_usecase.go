package usecase

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
)

type ServerUserUsecase interface {
	UserGetByID(context.Context, int64) (*entity.User, error)
	UserGetByIDList(context.Context, int64, int64) ([]entity.User, error)
	UserGetByName(context.Context, string) (*entity.User, error)
	Store(context.Context, *entity.User) (*entity.User, error)
}
