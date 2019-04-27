package repository

import (
	"context"
	"log"

	"github.com/NoahOrberg/evileye/entity"
	p2pclient "github.com/NoahOrberg/evileye/p2p/client"
	"github.com/jmoiron/sqlx"
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

		// get voted user id
		uids := make([]string, 0, sumofvote)
		rows, err := db.Query(`SELECT * FROM votes WHERE tarekomiid = ?`, t.ID)
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

			uids = append(uids, us.ScreenName)
		}

		u := new(entity.User)

		if err := db.Get(u, q, t.TargetUserID); err != nil {
			return false, err
		}

		ic.SentTxToLeaderNode(ctx, u.ScreenName, t.URL, t.Description, uids)
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
