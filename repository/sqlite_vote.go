package repository

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/log"
	p2pclient "github.com/NoahOrberg/evileye/p2p/client"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type SqliteVoteRepository struct {
	db *sqlx.DB
}

func NewSqliteVoteRepository(db *sqlx.DB) SqliteVoteRepository {
	return SqliteVoteRepository{db: db}
}

func checkVotes(ctx context.Context, tid int64, db *sqlx.DB, ic p2pclient.InternalClient) (bool, error) {
	qstr := `SELECT count(*) FROM votes WHERE tarekomiid = ?`

	t := &entity.Tarekomi{}

	res, err := db.Query(qstr, tid)
	if err != nil {
		return false, err
	}

	var sumofvote int64
	for res.Next() {
		if err := res.Scan(&sumofvote); err != nil {
			log.L().Error("when invoked checkVotes, happens an error: ", zap.Error(err))
			continue
		}
	}

	if t.Threshold <= sumofvote {
		log.L().Error("Tarekomi approved")

		t.Status = 1
		_, err := UpdateTarekomiState(ctx, *t, db)
		if err != nil {
			return false, err
		}

		// get voted user id
		unames := make([]string, 0, sumofvote)
		rows, err := db.Query(`SELECT userid FROM votes WHERE tarekomiid = ?`, t.ID)
		if err != nil {
			return false, err
		}

		q := `SELECT * FROM users WHERE id = ?`
		for rows.Next() {
			vu := entity.Vote{}
			if err := rows.Scan(
				&vu.UserID,
			); err != nil {
				return false, err
			}

			us := new(entity.User)
			if err := db.Get(us, q, vu.UserID); err != nil {
				return false, err
			}

			unames = append(unames, us.ScreenName)
		}

		u := new(entity.User)

		if err := db.Get(u, q, t.TargetUserID); err != nil {
			return false, err
		}

		log.L().Info("SentTxToLeaderNode", zap.String("username", u.ScreenName), zap.String("url", t.URL), zap.String("description", t.Description), zap.Strings("unames", unames))
		err = ic.SentTxToLeaderNode(ctx, u.ScreenName, t.URL, t.Description, unames)
		if err != nil {
			log.L().Error("SentTxToleaderNode error", zap.Error(err))
		}
	}

	return true, nil
}

func (r *SqliteVoteRepository) NewVoting(ctx context.Context, v *entity.Vote, ic p2pclient.InternalClient) error {
	qstr := `INSERT INTO votes(userid, tarekomiid, description) VALUES(?, ?, ?)`

	_, err := r.db.Exec(qstr, v.UserID, v.TarekomiID, v.Description)
	if err != nil {
		return err
	}

	_, err = checkVotes(ctx, v.TarekomiID, r.db, ic)
	if err != nil {
		return err
	}

	return nil
}
