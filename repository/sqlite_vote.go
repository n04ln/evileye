package repository

import (
	"context"
	"log"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/jmoiron/sqlx"
)

type SqliteVoteRepository struct {
	db *sqlx.DB
}

func NewSqliteVoteRepository(db *sqlx.DB) SqliteVoteRepository {
	return SqliteVoteRepository{db: db}
}

func checkVotes(ctx context.Context, tid int64, db *sqlx.DB) (bool, error) {
	qstr := `SELECT count(*) FROM votes WHERE tarekomiid = ?`

	t := &entity.Tarekomi{}

	res, err := db.Query(qstr, tid)
	if err != nil {
		return false, err
	}

	var sumofvote int64
	for res.Next() {
		if err := res.Scan(&sumofvote); err != nil {
			log.Printf("when invoked checkVotes, happens an error: %s", err)
			continue
		}
	}

	if t.Threshold >= sumofvote {
		log.Println("Tarekomi Approved!")
		t.Status = 1
		_, err := UpdateTarekomiState(ctx, *t, db)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (r *SqliteVoteRepository) NewVoting(ctx context.Context, v *entity.Vote) error {
	qstr := `INSERT INTO votes(userid, tarekomiid, description) VALUES(?, ?, ?)`

	_, err := r.db.Exec(qstr, v.UserID, v.TarekomiID, v.Description)
	if err != nil {
		return err
	}

	_, err = checkVotes(ctx, v.TarekomiID, r.db)
	if err != nil {
		return err
	}

	return nil
}
