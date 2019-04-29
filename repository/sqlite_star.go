package repository

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/jmoiron/sqlx"
)

type sqliteStarRepository struct {
	db *sqlx.DB
}

func NewSqliteStarRepository(db *sqlx.DB) StarRepository {
	return &sqliteStarRepository{db: db}
}

func (r *sqliteStarRepository) GetStaredTarekomiID(ctx context.Context, uid int64) ([]int64, error) {
	qstr := `SELECT tarekomiid FROM stars WHERE id = ?`

	rows, err := r.db.Query(qstr, uid)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0)

	for rows.Next() {
		s := entity.Star{}
		if err := rows.Scan(
			&s.TarekomiID,
		); err != nil {
			return nil, err
		}
		ids = append(ids, s.TarekomiID)
	}

	return ids, nil

}

func (r *sqliteStarRepository) DeleteStar(ctx context.Context, s *entity.Star) error {
	qstr := `DELETE FROM stars WHERE userid = ? AND tarekomiid = ?`

	_, err := r.db.Exec(qstr, s.UserID, s.TarekomiID)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteStarRepository) Store(ctx context.Context, s *entity.Star) (int64, error) {
	qstr := `INSERT INTO stars(userid, tarekomiid) VALUES(?, ?)`

	res, err := r.db.Exec(qstr, s.UserID, s.TarekomiID)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
