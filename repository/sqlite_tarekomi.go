package repository

import (
	"context"

	"github.com/NoahOrberg/evileye/entity"
	pb "github.com/NoahOrberg/evileye/protobuf"
	"github.com/jmoiron/sqlx"
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
	qstr := `UPDATE tarekomi SET state = ? WHERE id = ?`

	_, err := db.Exec(qstr, newtarekomi.Status, newtarekomi.ID)
	if err != nil {
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

func (r *sqliteTarekomiRepository) GetTarekomiBoard(ctx context.Context, limit, offset int64) (pb.TarekomiSummaries, error) {

	qstr := `SELECT * FROM tarekomi WHERE status = 0 ORDER BY id LIMIT ? OFFSET ?`

	tk := make([]*pb.TarekomiSummary, 0, limit)
	ts := new(pb.TarekomiSummaries)

	rows, err := r.db.Query(qstr, limit, offset)
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
