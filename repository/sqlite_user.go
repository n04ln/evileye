package repository

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/jmoiron/sqlx"
)

type sqliteUserRepository struct {
	db *sqlx.DB
}

func NewSqliteUserRepository(db *sqlx.DB) UserRepository {
	return &sqliteUserRepository{db}
}

func (r *sqliteUserRepository) UserGetByID(ctx context.Context, id int64) (*entity.User, error) {
	panic("not impl")
}

func (r *sqliteUserRepository) UserGetByIDList(ctx context.Context, ids []int64) ([]entity.User, error) {
	panic("not impl")
}

func (r *sqliteUserRepository) UserGetByName(ctx context.Context, uname string) (*entity.User, error) {

	qstr := `SELECT * FROM user WHERE screenname = ?`
	u := new(entity.User)

	if err := r.db.Get(u, qstr, uname); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *sqliteUserRepository) Store(ctx context.Context, usr *entity.User) error {
	panic("not impl")
}
