package repository

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
)

type StarRepository interface {
	GetStaredTarekomiID(context.Context, int64) ([]int64, error)
	DeleteStar(context.Context, *entity.Star) error
	Store(context.Context, *entity.Star) (int64, error)
}
