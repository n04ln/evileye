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
	qstr := `SELECT * FROM user WHERE id = ?`
	u := new(entity.User)

	if err := r.db.Get(u, qstr, id); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *sqliteUserRepository) UserGetByIDList(ctx context.Context, limit, offset int64) ([]entity.User, error) {
	qstr := `SELECT * FROM user ORDER By id LIMIT ? OFFSET ?`
	us := make([]entity.User, 0, limit)

	rows, err := r.db.Query(qstr, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := entity.User{}
		if err := rows.Scan(
			&u.ID,
			&u.ScreenName,
		); err != nil {
			return nil, err
		}
		us = append(us, u)

	}

	return us, nil
}

func (r *sqliteUserRepository) UserGetByName(ctx context.Context, uname string) (*entity.User, error) {

	qstr := `SELECT * FROM user WHERE screenname = ?`
	u := new(entity.User)

	if err := r.db.Get(u, qstr, uname); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *sqliteUserRepository) Store(ctx context.Context, usr *entity.User) (*entity.User, error) {
	qstr := `INSERT INTO users(screenname, password) VALUES(?, ?)`
	res, err := r.db.Exec(qstr, usr.ScreenName, usr.Password)
	if err != nil {
		return nil, err
	}

	usr.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return usr, nil
}
