package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/NoahOrberg/evileye/entity"
	"github.com/NoahOrberg/evileye/interceptor"
	"github.com/NoahOrberg/evileye/log"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	PENDING  = 0
	APPROVED = 1
	REJECTED = 2
)

type sqliteTarekomiRepository struct {
	db *sqlx.DB
}

func NewSqliteTarekomiRepository(db *sqlx.DB) TarekomiRepository {
	return &sqliteTarekomiRepository{db}
}

func UpdateTarekomiState(ctx context.Context, newtarekomi entity.Tarekomi, db *sqlx.DB) (entity.Tarekomi, error) {
	qstr := `UPDATE tarekomi SET status = ? WHERE id = ?`

	_, err := db.Exec(qstr, newtarekomi.Status, newtarekomi.ID)
	if err != nil {
		log.L().Error("exec sql failed",
			zap.String("q", qstr),
			zap.Any("args", []interface{}{newtarekomi.Status, newtarekomi.ID}),
			zap.Error(err))
		return newtarekomi, err
	}

	return newtarekomi, nil
}

func getUserByID(ctx context.Context, uid int64, db *sqlx.DB) (*entity.User, error) {
	qstr := `SELECT * FROM users WHERE id = ?`
	u := new(entity.User)

	if err := db.Get(u, qstr, uid); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *sqliteTarekomiRepository) GetTarekomiFromUser(ctx context.Context, id, limit, offset int64) (pb.TarekomiSummaries, error) {

	tk := make([]*pb.TarekomiSummary, 0, limit)
	ts := new(pb.TarekomiSummaries)

	u, err := getUserByID(ctx, id, r.db)
	if err != nil {
		return *ts, err
	}

	qstr := `SELECT * FROM tarekomi WHERE id = ? LIMIT ? OFFSET ?`

	rows, err := r.db.Query(qstr, id, limit, offset)
	if err != nil {
		return *ts, err
	}

	for rows.Next() {
		t := entity.Tarekomi{}

		if err := rows.Scan(
			&t.ID,
			&t.Status,
			&t.Threshold,
			&t.TargetUserID,
			&t.URL,
			&t.Description,
		); err != nil {
			return *ts, err
		}

		u1, err := getUserByID(ctx, t.ID, r.db)
		if err != nil {
			return *ts, err
		}

		tp := pb.TarekomiSummary{
			Tarekomi: &pb.Tarekomi{
				TargetUserName: u1.ScreenName,
				Url:            t.URL,
				Desc:           t.Description,
			},
			UserName: u.ScreenName,
		}

		tk = append(tk, &tp)
	}

	ts.Tarekomis = tk

	return *ts, nil
}

func (r *sqliteTarekomiRepository) GetTarekomiApproved(ctx context.Context, uid int64) ([]*pb.Tarekomi, error) {

	ts := make([]*pb.Tarekomi, 0)

	qstr := `SELECT * FROM tarekomi WHERE targetuserid = ? AND status = 1`

	rows, err := r.db.Query(qstr, uid)
	if err != nil {
		return ts, err
	}

	for rows.Next() {
		t := entity.Tarekomi{}
		if err := rows.Scan(
			&t.ID,
			&t.Status,
			&t.Threshold,
			&t.TargetUserID,
			&t.URL,
			&t.Description,
		); err != nil {
			return ts, err
		}

		u1, err := getUserByID(ctx, t.ID, r.db)
		if err != nil {
			return ts, err
		}

		tp := &pb.Tarekomi{
			TargetUserName: u1.ScreenName,
			Url:            t.URL,
			Desc:           t.Description,
		}

		ts = append(ts, tp)
	}

	return ts, nil
}

func (r *sqliteTarekomiRepository) GetTarekomiBoard(ctx context.Context, limit, offset int64) (pb.TarekomiSummaries, error) {

	tk := make([]*pb.TarekomiSummary, 0, limit)
	ts := new(pb.TarekomiSummaries)

	ui := interceptor.GetUserMetaData(ctx)

	voted, err := VotedFromUserID(ctx, ui.ID, r.db)
	if err != nil {
		log.L().Error("VotedFromUserID error", zap.Error(err))
		return *ts, err
	}

	var rows *sql.Rows

	if len(voted) > 0 {
		qstr := `SELECT * FROM tarekomi WHERE status = 0 
		ORDER BY id LIMIT ? OFFSET ? NOT IN (? ` + strings.Repeat(`, ?`, len(voted)-1) + `)`

		args := make([]interface{}, 0, len(voted)+2)
		args = append(args, limit)
		args = append(args, offset)
		for _, d := range voted {
			args = append(args, d)
		}

		log.L().Info("exec query and using evil solution", zap.String("qstr", qstr), zap.Any("args", args))
		rows, err = r.db.Query(qstr, args...)
	} else {
		qstr := `SELECT * FROM tarekomi WHERE status = 0 ORDER BY id LIMIT ? OFFSET ?`
		rows, err = r.db.Query(qstr, limit, offset)
	}

	if err != nil {
		return *ts, err
	}

	for rows.Next() {
		t := entity.Tarekomi{}

		if err := rows.Scan(
			&t.ID,
			&t.Status,
			&t.Threshold,
			&t.TargetUserID,
			&t.URL,
			&t.Description,
		); err != nil {
			return *ts, err
		}

		n, err := getUserByID(ctx, t.TargetUserID, r.db)
		if err != nil {
			return *ts, err
		}

		tp := pb.TarekomiSummary{
			Tarekomi: &pb.Tarekomi{
				TargetUserName: n.ScreenName,
				Url:            t.URL,
				Desc:           t.Description,
			},
			UserName: n.ScreenName,
		}

		tk = append(tk, &tp)
	}

	ts.Tarekomis = tk

	return *ts, nil
}

// tarekomiを登録(投票街になるのでstatusはデフォで0)
func (r *sqliteTarekomiRepository) Store(ctx context.Context, t entity.Tarekomi) (int64, error) {
	qstr := `INSERT INTO tarekomi(status, threshold, targetuserid, url, description) VALUES(?, ?, ?, ?, ?)`

	res, err := r.db.Exec(qstr, PENDING, t.Threshold, t.TargetUserID, t.URL, t.Description)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *sqliteTarekomiRepository) UpdateTarekomiState(ctx context.Context, newtarekomi entity.Tarekomi) (entity.Tarekomi, error) {
	qstr := `UPDATE tarekomi SET state = ? WHERE id = ?`

	_, err := r.db.Exec(qstr, newtarekomi.Status, newtarekomi.ID)
	if err != nil {
		return newtarekomi, err
	}

	return newtarekomi, nil
}

func (r *sqliteTarekomiRepository) GetTarekomiFromID(ctx context.Context, tid int64) (pb.TarekomiSummary, error) {
	qstr := `SELECT * FROM tarekomi WHERE id = ?`

	rows, err := r.db.Query(qstr, tid)
	if err != nil {
		return pb.TarekomiSummary{}, err
	}

	tp := pb.TarekomiSummary{}

	for rows.Next() {
		t := entity.Tarekomi{}

		if err := rows.Scan(
			&t.ID,
			&t.Status,
			&t.Threshold,
			&t.TargetUserID,
			&t.URL,
			&t.Description,
		); err != nil {
			return pb.TarekomiSummary{}, err
		}

		n, err := getUserByID(ctx, t.TargetUserID, r.db)
		if err != nil {
			return pb.TarekomiSummary{}, err
		}

		tp = pb.TarekomiSummary{
			Tarekomi: &pb.Tarekomi{
				TargetUserName: n.ScreenName,
				Url:            t.URL,
				Desc:           t.Description,
			},
			UserName: n.ScreenName,
		}
	}

	return tp, nil
}
