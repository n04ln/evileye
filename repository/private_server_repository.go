package repository

import "context"

type PrivateServerRepository interface {
	UserGetByID(ctx context.Context, id int64)
}
